package gbt19056

// ExternalPowerLog ..
type ExternalPowerLog struct {
	dataBlockMeta
	Records []ExternalPowerLogRecord `json:"records"`
}

// ExternalPowerLogRecord ...
type ExternalPowerLogRecord struct {
	Ts   DateTime `json:"ts,string"`
	Type HexUint8 `json:"type"` // 0x01: pluged; 0x02: unpluged;
}

// DumpData ExternalPowerLog
func (e *ExternalPowerLog) DumpData() ([]byte, error) {
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

// DumpData ExternalPowerLogRecord
func (e *ExternalPowerLogRecord) DumpData() ([]byte, error) {
	var err error

	var ts []byte
	ts, err = e.Ts.DumpData()

	buff := append(ts, uint8(e.Type))

	return buff, err
}
