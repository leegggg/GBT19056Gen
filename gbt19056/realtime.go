package gbt19056

// RealTime ..
type RealTime struct {
	dataBlockMeta
	Now DateTime `json:"now,string"`
}

// DumpData RealTime
func (e *RealTime) DumpData() ([]byte, error) {
	bs, err := e.Now.DumpData()
	bs, err = e.linkDumpedData(bs)
	return bs, err
}

// LoadBinary RealTime Table A.8, Code 0x00
func (e *RealTime) LoadBinary(buffer []byte, meta dataBlockMeta) {
	e.dataBlockMeta = meta
	e.Now.LoadBinary(buffer[0:6])
	return
}
