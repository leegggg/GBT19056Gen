package gbt19056

// OvertimeLog ..
type OvertimeLog struct {
	dataBlockMeta
	Records []OvertimeLogRecord `json:"records"`
}

// OvertimeLogRecord ...
type OvertimeLogRecord struct {
	DriverID      string   `json:"driver_id"`
	Start         DateTime `json:"start,string"`
	End           DateTime `json:"end,string"`
	PositionStart Position `json:"position_start"`
	PositionEnd   Position `json:"position_end"`
}

// DumpData SpeedLog
func (e *OvertimeLog) DumpData() ([]byte, error) {
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

// DumpData OvertimeLogRecord
func (e *OvertimeLogRecord) DumpData() ([]byte, error) {
	var err error

	var start, end, startPost, endPost []byte
	driverID := make([]byte, 18)
	// TODO should check length. ASCII
	copy(driverID, []byte(e.DriverID))

	start, err = e.Start.DumpData()
	end, err = e.End.DumpData()
	startPost, err = e.PositionStart.DumpData()
	endPost, err = e.PositionEnd.DumpData()

	buff := append(driverID, start...)
	buff = append(buff, end...)
	buff = append(buff, startPost...)
	buff = append(buff, endPost...)

	return buff, err
}

// LengthOvertimeLogRecord ...
const LengthOvertimeLogRecord = 50

// LoadBinary SpeedLog Table A.16, Code 0x08
func (e *OvertimeLog) LoadBinary(buffer []byte, meta dataBlockMeta) {
	dataLength := LengthOvertimeLogRecord
	e.dataBlockMeta = meta
	for ptr := 0; ptr < len(buffer); ptr = ptr + dataLength {
		record := new(OvertimeLogRecord)
		record.LoadBinary(buffer[ptr : ptr+dataLength])
		e.Records = append(e.Records, *record)
	}
	return
}

// LoadBinary DriverLogRecord Table A.25
func (e *OvertimeLogRecord) LoadBinary(buffer []byte) {
	e.DriverID = bytesToStr(buffer[0:18])
	e.Start.LoadBinary(buffer[18:24])
	e.End.LoadBinary(buffer[24:30])
	e.PositionStart.LoadBinary(buffer[30:40])
	e.PositionEnd.LoadBinary(buffer[40:50])
	return
}
