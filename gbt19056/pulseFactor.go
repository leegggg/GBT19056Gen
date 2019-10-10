package gbt19056

import (
	"encoding/binary"
)

// PulseFactor ..
type PulseFactor struct {
	dataBlockMeta
	Now    DateTime `json:"now,string"`
	Factor uint16   `json:"Factor"`
}

// DumpData PulseFactor
func (e *PulseFactor) DumpData() ([]byte, error) {
	now, err := e.Now.DumpData()

	factor := make([]byte, 2)
	binary.BigEndian.PutUint16(factor, e.Factor)

	bs := append(now, factor...)

	bs, err = e.linkDumpedData(bs)
	return bs, err
}

// LoadBinary PulseFactor Table A.10, Code 0x04
func (e *PulseFactor) LoadBinary(buffer []byte, meta dataBlockMeta) {
	e.dataBlockMeta = meta
	e.Now.LoadBinary(buffer[0:6])
	e.Factor = binary.BigEndian.Uint16(buffer[6:8])
	return
}
