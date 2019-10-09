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
