package gbt19056

import (
	"fmt"
)

// LengthPositionLogRecord ...
const LengthPositionLogRecord = 666

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

// LoadBinary SpeedLog Table A.16, Code 0x08
func (e *PositionLog) LoadBinary(buffer []byte, meta dataBlockMeta) {
	e.dataBlockMeta = meta
	for ptr := 0; ptr < len(buffer); ptr = ptr + LengthPositionLogRecord {
		record := new(PositionLogRecord)
		record.LoadBinary(buffer[ptr : ptr+LengthPositionLogRecord])
		e.Records = append(e.Records, *record)
	}
	return
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

	// Table A.19, Full the block with 0xFF if length is not LengthPositionLogRecord
	blockLength := LengthPositionLogRecord
	if len(buff) != blockLength {
		err = error(fmt.Errorf("buffer size of SpeedLogRecord is not %d", blockLength))
		for i := len(buff); i < blockLength; i++ {
			buff = append(buff, 0xFF)
		}
	}
	return buff, err
}

// LoadBinary SpeedLogRecord Table A.17
func (e *PositionLogRecord) LoadBinary(buffer []byte) {
	lengthPositionWithSpeed := 11
	e.Start.LoadBinary(buffer[0:6])
	for ptr := 6; ptr < len(buffer); ptr = ptr + lengthPositionWithSpeed {
		point := new(positionWithSpeed)
		point.LoadBinary(buffer[ptr : ptr+lengthPositionWithSpeed])
		e.Positions = append(e.Positions, *point)
	}
	return
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

// LoadBinary SpeedLogRecord Table A.17
func (e *positionWithSpeed) LoadBinary(buffer []byte) {
	e.Speed = buffer[10]
	e.Position.LoadBinary(buffer[0:10])
	return
}
