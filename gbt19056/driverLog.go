package gbt19056

// DriverLog ..
type DriverLog struct {
	dataBlockMeta
	Records []DriverLogRecord `json:"records"`
}

// DriverLogRecord ...
type DriverLogRecord struct {
	Ts       DateTime `json:"ts,string"`
	DriverID string   `json:"driver_id"`
	Type     HexUint8 `json:"type"` // 0x01: Login; 0x02: Logout; Other resversed
}

// DumpData DriverLog
func (e *DriverLog) DumpData() ([]byte, error) {
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

// DumpData DriverLogRecord
func (e *DriverLogRecord) DumpData() ([]byte, error) {
	var err error

	driverID := make([]byte, 18)
	// TODO should check length. ASCII
	copy(driverID, []byte(e.DriverID))

	var ts []byte
	ts, err = e.Ts.DumpData()

	buff := append(ts, driverID...)
	buff = append(buff, uint8(e.Type))

	return buff, err
}

// LengthDriverLogRecord ...
const LengthDriverLogRecord = 25

// LoadBinary SpeedLog Table A.16, Code 0x12
func (e *DriverLog) LoadBinary(buffer []byte, meta dataBlockMeta) {
	e.dataBlockMeta = meta
	for ptr := 0; ptr < len(buffer); ptr = ptr + LengthDriverLogRecord {
		record := new(DriverLogRecord)
		record.LoadBinary(buffer[ptr : ptr+LengthDriverLogRecord])
		e.Records = append(e.Records, *record)
	}
	return
}

// LoadBinary DriverLogRecord Table A.25
func (e *DriverLogRecord) LoadBinary(buffer []byte) {
	e.Ts.LoadBinary(buffer[0:6])
	e.DriverID = bytesToStr(buffer[6:24])
	e.Type = HexUint8(buffer[24])
	return
}
