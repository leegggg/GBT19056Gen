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

// LoadBinary SpeedLog Table A.16, Code 0x08
func (e *ConfigChangeLog) LoadBinary(buffer []byte, meta dataBlockMeta) {
	dataLength := 7
	e.dataBlockMeta = meta
	for ptr := 0; ptr < len(buffer); ptr = ptr + dataLength {
		record := new(ConfigChangeLogRecord)
		record.LoadBinary(buffer[ptr : ptr+dataLength])
		e.Records = append(e.Records, *record)
	}
	return
}

// LoadBinary DriverLogRecord Table A.25
func (e *ConfigChangeLogRecord) LoadBinary(buffer []byte) {
	e.Ts.LoadBinary(buffer[0:6])
	e.Type = HexUint8(buffer[6])
	return
}
