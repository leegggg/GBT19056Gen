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
