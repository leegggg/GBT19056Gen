package gbt19056

// OvertimeLog ..
type OvertimeLog struct {
	dataBlockMeta
	Records []OvertimeLogRecord `json:"records"`
}

// OvertimeLogRecord ...
type OvertimeLogRecord struct {
	Start         DateTime `json:"start,string"`
	End           DateTime `json:"end,string"`
	PositionStart Position `json:"position_start"`
	PositionEnd   Position `json:"position_end"`
}
