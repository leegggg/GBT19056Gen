package gbt19056

// PositionLog ..
type PositionLog struct {
	dataBlockMeta
	Records []PositionLogRecord `json:"records"`
}

// PositionLogRecord ...
type PositionLogRecord struct {
	Start     string `json:"start"`
	Positions []struct {
		Speed    uint8    `json:"speed"`
		Position Position `json:"position"`
	} `json:"positions"`
}
