package gbt19056

// AccidentLog ..
type AccidentLog struct {
	dataBlockMeta
	Records []AccidentLogRecord `json:"records"`
}

// AccidentLogRecord ...
type AccidentLogRecord struct {
	Ts            DateTime      `json:"ts,string"`
	EndPosition   Position      `json:"position"`
	SpeedStatuses []SpeedStatus `json:"speed_statuses"`
}
