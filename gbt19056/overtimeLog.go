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
