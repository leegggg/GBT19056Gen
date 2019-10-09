package gbt19056

import "encoding/binary"

// RecoderID ..
type RecoderID struct {
	dataBlockMeta
	CCC     string   `json:"CCC"`
	Version string   `json:"version"`
	Dop     DateTime `json:"dop,string"`
	Sn      uint32   `json:"sn"`
	Comment string   `json:"comment"`
}

// DumpData RecoderID
func (e *RecoderID) DumpData() ([]byte, error) {
	var err error

	// ASCII ID should be safe to be copy directly
	ccc := make([]byte, 7)
	copy(ccc, []byte(e.CCC))

	// TODO: Check length
	version := make([]byte, 16)
	copy(version, []byte(e.Version))

	var dop []byte
	dop, err = e.Dop.DumpDataShort()

	sn := make([]byte, 4)
	binary.BigEndian.PutUint32(sn, e.Sn)

	// TODO： Unfinished comment logic
	comment := make([]byte, 5)
	copy(comment, []byte(e.Comment))

	bs := append(ccc, version...)
	bs = append(bs, dop...)
	bs = append(bs, sn...)
	bs = append(bs, comment...)

	bs, err = e.linkDumpedData(bs)
	return bs, err
}
