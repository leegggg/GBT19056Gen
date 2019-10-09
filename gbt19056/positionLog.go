package gbt19056

import (
	"fmt"
)

// PositionLog ..
type PositionLog struct {
	dataBlockMeta
	Records []PositionLogRecord `json:"records"`
}

// PositionLogRecord ...
type PositionLogRecord struct {
	Start     DateTime            `json:"start,string"`
	Positions []positionWithSpeed `json:"positions"`
}

type positionWithSpeed struct {
	Speed    uint8    `json:"speed"`
	Position Position `json:"position"`
}

// DumpData PositionLog
func (e *PositionLog) DumpData() ([]byte, error) {
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
func (e *PositionLogRecord) DumpData() ([]byte, error) {
	var err error
	buff := []byte{}
	if len(e.Positions) <= 0 {
		return buff, err
	}
	buff, err = e.Start.DumpData()
	// Table A.19, Should have 60 points(min) in the block, if there are not enough copy last one
	var last []byte
	last, err = (e.Positions[0]).dumpData()
	for i := 0; i < 60; i++ {
		if i < len(e.Positions) {
			last, err = (e.Positions[i]).dumpData()
		}
		buff = append(buff, last...)
	}

	// Table A.19, Full the block with 0xFF if length is not 126
	blockLength := 666
	if len(buff) != blockLength {
		err = error(fmt.Errorf("buffer size of SpeedLogRecord is not %d", blockLength))
		for i := len(buff); i < blockLength; i++ {
			buff = append(buff, 0xFF)
		}
	}
	return buff, err
}

// dumpData positionWithSpeed
func (e *positionWithSpeed) dumpData() ([]byte, error) {
	var err error
	var buff []byte

	var position []byte
	position, err = e.Position.DumpData()

	buff = append(position, e.Speed)
	return buff, err
}
