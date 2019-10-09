package gbt19056

// SpeedLog ..
type SpeedLog struct {
	dataBlockMeta
	Records []SpeedLogRecord `json:"records"`
}

// SpeedLogRecord ...
type SpeedLogRecord struct {
	Start         DateTime      `json:"start,string"`
	SpeedStatuses []SpeedStatus `json:"speed_statuses"`
}
