package gbt19056

// StatusSignalConfig ..
type StatusSignalConfig struct {
	dataBlockMeta
	Now    DateTime   `json:"now,string"`
	Status Status     `json:"status"`
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
