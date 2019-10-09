package gbt19056

import "encoding/binary"

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
