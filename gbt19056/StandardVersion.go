package gbt19056

import (
	"encoding/binary"

	"github.com/leegggg/GBT19056Gen/utils/bcd"
)

// StandardVersion ...
type StandardVersion struct {
	dataBlockMeta
	Year   uint8    `json:"year"`
	Update HexUint8 `json:"update"`
}

// DumpDate ...
func (e *StandardVersion) DumpDate() ([]byte, error) {
	meta, err := e.dataBlockMeta.DumpDate()

	bs := make([]byte, 2)
	bs[0] = bcd.FromUint8(e.Year)
	bs[1] = (byte)(e.Update)

	// Data Length
	dataLength := make([]byte, 4)
	binary.BigEndian.PutUint32(dataLength, (uint32)(len(bs)))
	meta = append(meta, dataLength...)
	bs = append(meta, bs...)
	return bs, err
}
