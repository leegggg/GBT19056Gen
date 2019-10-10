package gbt19056

// DriverInfo ..
type DriverInfo struct {
	dataBlockMeta
	ID string `json:"id"`
}

// DumpData StandardVersion
func (e *DriverInfo) DumpData() ([]byte, error) {

	bs := make([]byte, 18)
	// ASCII ID should be safe to be copy directly
	copy(bs, []byte(e.ID))
	bs, _ = e.linkDumpedData(bs)
	return bs, nil
}

// LoadBinary DriverInfo Table A.6, Code 0x00
func (e *DriverInfo) LoadBinary(buffer []byte, meta dataBlockMeta) {
	e.dataBlockMeta = meta
	e.ID = bytesToStr(buffer[0:18])
	return
}
