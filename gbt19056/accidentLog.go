package gbt19056

import "fmt"

// AccidentLog ..
type AccidentLog struct {
	dataBlockMeta
	Records []AccidentLogRecord `json:"records"`
}

// AccidentLogRecord ...
type AccidentLogRecord struct {
	Ts            DateTime      `json:"ts,string"`
	DriverID      string        `json:"driver_id"`
	EndPosition   Position      `json:"position"`
	SpeedStatuses []SpeedStatus `json:"speed_statuses"`
}

// DumpData PositionLog
func (e *AccidentLog) DumpData() ([]byte, error) {
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

// DumpData SpeedLogRecord
func (e *AccidentLogRecord) DumpData() ([]byte, error) {
	var err error
	buff := []byte{}
	buff, err = e.Ts.DumpData()

	driverID := make([]byte, 18)
	// TODO should check length. ASCII
	copy(driverID, []byte(e.DriverID))
	buff = append(buff, driverID...)

	if len(e.SpeedStatuses) > 0 {
		// Table A.22, Should have 100 points(5pps in 20s) in the block, if there are not enough copy last one
		nbPoints := 100
		var last []byte
		last, err = (e.SpeedStatuses[0]).DumpData()
		for i := 0; i < nbPoints; i++ {
			if i < len(e.SpeedStatuses) {
				last, err = (e.SpeedStatuses[i]).DumpData()
			}
			buff = append(buff, last...)
		}
	} else {
		err = error(fmt.Errorf("Got empty speed status in accident data please check input data"))
	}

	var position []byte
	position, err = e.EndPosition.DumpData()
	buff = append(buff, position...)

	// Table A.22, Full the block with 0xFF if length is not 234
	blockLength := 234
	if len(buff) != blockLength {
		err = error(fmt.Errorf("buffer size of SpeedLogRecord is not %d, Table A.22", blockLength))
		for i := len(buff); i < blockLength; i++ {
			buff = append(buff, 0xFF)
		}
	}
	return buff, err
}
