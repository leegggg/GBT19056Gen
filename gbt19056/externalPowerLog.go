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

// LoadBinary SpeedLog Table A.16, Code 0x08
func (e *ExternalPowerLog) LoadBinary(buffer []byte, meta dataBlockMeta) {
	dataLength := 7
	e.dataBlockMeta = meta
	for ptr := 0; ptr < len(buffer); ptr = ptr + dataLength {
		record := new(ExternalPowerLogRecord)
		record.LoadBinary(buffer[ptr : ptr+dataLength])
		e.Records = append(e.Records, *record)
	}
	return
}

// LoadBinary DriverLogRecord Table A.25
func (e *ExternalPowerLogRecord) LoadBinary(buffer []byte) {
	e.Ts.LoadBinary(buffer[0:6])
	e.Type = HexUint8(buffer[6])
	return
}
