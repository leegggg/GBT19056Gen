package gbt19056

import "fmt"

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
