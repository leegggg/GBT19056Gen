package gbt19056

import (
	"github.com/leegggg/GBT19056Gen/utils/bcd"
)

// StandardVersion ...
type StandardVersion struct {
	dataBlockMeta
	Year   uint8    `json:"year"`
	Update HexUint8 `json:"update"`
}

// DumpData StandardVersion
func (e *StandardVersion) DumpData() ([]byte, error) {
	var err error
	bs := make([]byte, 2)
	bs[0] = bcd.FromUint8(e.Year)
	bs[1] = (byte)(e.Update)
	bs, err = e.linkDumpedData(bs)
	return bs, err
}

// LoadBinary StandardVersion Table A.6, Code 0x00
func (e *StandardVersion) LoadBinary(buffer []byte, meta dataBlockMeta) {
	e.dataBlockMeta = meta
	e.Year = bcd.ToUint8(buffer[0])
	e.Update = HexUint8(buffer[1])
	return
}
