package gbt19056

// LengthStatusName ...
const LengthStatusName = 10

// StatusSignalConfig ..
type StatusSignalConfig struct {
	dataBlockMeta
	Now    DateTime   `json:"now,string"`
	Config [][]string `json:"config"`
}

// DumpData StatusSignalConfig
func (e *StatusSignalConfig) DumpData() ([]byte, error) {
	var buff []byte
	var err error

	var now []byte
	now, err = e.Now.DumpData()

	// TODO: Check length
	status := uint8(len(e.Config))

	bs := append(now, status)

	for _, c := range e.Config {
		for _, v := range c {
			buff, err = EncodeGBK(([]byte)(v))
			sub := make([]byte, 10)
			copy(sub, buff)
			bs = append(bs, sub...)
		}
	}

	bs, err = e.linkDumpedData(bs)
	return bs, err
}

// LoadBinary StatusSignalConfig Table A.12, Code 0x06
func (e *StatusSignalConfig) LoadBinary(buffer []byte, meta dataBlockMeta) {
	e.dataBlockMeta = meta
	e.Now.LoadBinary(buffer[0:6])
	nbStatus := int(buffer[6])
	ptr := 7
	for i := 0; i < nbStatus; i++ {
		names := make([]string, 8)
		for j := 0; j < 8; j++ {
			names[j], _ = DecodeGBKStr(buffer[ptr : ptr+LengthStatusName])
			ptr += LengthStatusName
		}
		e.Config = append(e.Config, names)
	}
	return
}
