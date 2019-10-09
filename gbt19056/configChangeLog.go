package gbt19056

// ConfigChangeLog ..
type ConfigChangeLog struct {
	dataBlockMeta
	Records []ConfigChangeLogRecord `json:"records"`
}

// ConfigChangeLogRecord ...
type ConfigChangeLogRecord struct {
	Ts   DateTime `json:"ts,string"`
	Type HexUint8 `json:"event"`
}

// DumpData ConfigChangeLog
func (e *ConfigChangeLog) DumpData() ([]byte, error) {
	var err error
	var record []byte

	buff := []byte{}

	for _, v := range e.Records {
		record, err = v.DumpData()
		buff = append(buff, record...)
	}

	buff, err = e.linkDumpedData(buff)
	return buff, err
}

// DumpData ConfigChangeLogRecord
func (e *ConfigChangeLogRecord) DumpData() ([]byte, error) {
	var err error

	var ts []byte
	ts, err = e.Ts.DumpData()

	buff := append(ts, uint8(e.Type))

	return buff, err
}
