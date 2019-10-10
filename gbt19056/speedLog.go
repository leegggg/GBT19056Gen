package gbt19056

import "fmt"

// LengthSpeedLogRecord ...
const LengthSpeedLogRecord = 126

// SpeedLog ..
type SpeedLog struct {
	dataBlockMeta
	Records []SpeedLogRecord `json:"records"`
}

// SpeedLogRecord ...
type SpeedLogRecord struct {
	Start         DateTime      `json:"start,string"`
	SpeedStatuses []SpeedStatus `json:"speed_statuses"`
}

// DumpData SpeedLog
func (e *SpeedLog) DumpData() ([]byte, error) {
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

// LoadBinary SpeedLog Table A.16, Code 0x08
func (e *SpeedLog) LoadBinary(buffer []byte, meta dataBlockMeta) {
	e.dataBlockMeta = meta
	for ptr := 0; ptr < len(buffer); ptr = ptr + LengthSpeedLogRecord {
		record := new(SpeedLogRecord)
		record.LoadBinary(buffer[ptr : ptr+LengthSpeedLogRecord])
		e.Records = append(e.Records, *record)
	}
	return
}

// DumpData SpeedLogRecord
func (e *SpeedLogRecord) DumpData() ([]byte, error) {
	var err error
	buff := []byte{}
	if len(e.SpeedStatuses) <= 0 {
		return buff, err
	}
	buff, err = e.Start.DumpData()
	// Table A.17, Should have 60 points(sec) in the block, if there are not enough copy last one
	var last []byte
	last, err = (e.SpeedStatuses[0]).DumpData()
	for i := 0; i < 60; i++ {
		if i < len(e.SpeedStatuses) {
			last, err = (e.SpeedStatuses[i]).DumpData()
		}
		buff = append(buff, last...)
	}

	// Table A.17, Full the block with 0xFF if length is not 126
	if len(buff) != 126 {
		err = error(fmt.Errorf("buffer size of SpeedLogRecord is not 126"))
		for i := len(buff); i < 126; i++ {
			buff = append(buff, 0xFF)
		}
	}
	return buff, err
}

// LoadBinary SpeedLogRecord Table A.17
func (e *SpeedLogRecord) LoadBinary(buffer []byte) {
	e.Start.LoadBinary(buffer[0:6])
	for ptr := 6; ptr < len(buffer); ptr = ptr + LengthSpeedStatus {
		speed := new(SpeedStatus)
		speed.LoadBinary(buffer[ptr : ptr+LengthSpeedStatus])
		e.SpeedStatuses = append(e.SpeedStatuses, *speed)
	}
	return
}
