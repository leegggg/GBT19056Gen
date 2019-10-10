package gbt19056

import "fmt"

// SpeedStatusLog ..
type SpeedStatusLog struct {
	dataBlockMeta
	Records []SpeedStatusLogRecord `json:"records"`
}

// SpeedStatusLogRecord ...
type SpeedStatusLogRecord struct {
	Status     HexUint8 `json:"status"` // 0x01: normal; 0x02: error;
	Start      DateTime `json:"start,string"`
	End        DateTime `json:"end,string"`
	Speeds     []int    `json:"speeds"`
	SpeedsGNSS []int    `json:"speeds_gnss"`
}

// DumpData SpeedStatusLog
func (e *SpeedStatusLog) DumpData() ([]byte, error) {
	var err error
	var record []byte

	buff := []byte{}

	for _, v := range e.Records {
		record, err = v.DumpData()
		buff = append(buff, record...)
	}

	buff, err = e.linkDumpedData(buff)
	return buff, err
}

// DumpData SpeedStatusLogRecord
func (e *SpeedStatusLogRecord) DumpData() ([]byte, error) {
	var err error

	var start, end []byte

	buff := []uint8{uint8(e.Status)}

	start, err = e.Start.DumpData()
	buff = append(buff, start...)

	end, err = e.End.DumpData()
	buff = append(buff, end...)

	// Table A.32, Should have 60 points(sec) in the block, if there are not enough copy last one
	var lastSpeed, lastSpeedGNSS uint8
	lastSpeed = 0x00
	lastSpeedGNSS = 0x00
	for i := 0; i < 60; i++ {
		if i < len(e.Speeds) {
			lastSpeed = uint8(e.Speeds[i])
		}
		if i < len(e.SpeedsGNSS) {
			lastSpeedGNSS = uint8(e.SpeedsGNSS[i])
		}
		buff = append(buff, lastSpeed)
		buff = append(buff, lastSpeedGNSS)
	}

	// Table A.17, Full the block with 0xFF if length is not 126
	lengthBlock := LengthSpeedStatusLogRecord
	if len(buff) != lengthBlock {
		err = error(fmt.Errorf("buffer size of SpeedLogRecord is not %d Table A.33", lengthBlock))
		for i := len(buff); i < lengthBlock; i++ {
			buff = append(buff, 0xFF)
		}
	}
	return buff, err
}

// LengthSpeedStatusLogRecord ...
const LengthSpeedStatusLogRecord = 133

// LoadBinary SpeedLog Table A.16, Code 0x08
func (e *SpeedStatusLog) LoadBinary(buffer []byte, meta dataBlockMeta) {
	dataLength := LengthSpeedStatusLogRecord
	e.dataBlockMeta = meta
	for ptr := 0; ptr < len(buffer); ptr = ptr + dataLength {
		record := new(SpeedStatusLogRecord)
		record.LoadBinary(buffer[ptr : ptr+dataLength])
		e.Records = append(e.Records, *record)
	}
	return
}

// LoadBinary DriverLogRecord Table A.25
func (e *SpeedStatusLogRecord) LoadBinary(buffer []byte) {
	e.Status = HexUint8(buffer[0])
	e.Start.LoadBinary(buffer[1:7])
	e.End.LoadBinary(buffer[7:13])
	for ptr := 13; ptr < len(buffer); ptr = ptr + 2 {
		e.Speeds = append(e.Speeds, int(buffer[ptr]))
		e.SpeedsGNSS = append(e.SpeedsGNSS, int(buffer[ptr+1]))
	}
	return
}

// UnmarshalJSON ...
// func (s *SpeedStatusLog) UnmarshalJSON(data []byte) error {
// 	type Alias SpeedStatusLog
// 	aux := &struct {
// 		*Alias
// 	}{
// 		Alias: (*Alias)(s),
// 	}
// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}
// 	// s = SpeedStatusLog(*aux)
// 	// s.SomeCustomType = time.Unix(aux.SomeCustomType, 0)
// 	return nil
// }
// 	return nil
// }
// }
// 	// s.SomeCustomType = time.Unix(aux.SomeCustomType, 0)
// 	return nil
// }
