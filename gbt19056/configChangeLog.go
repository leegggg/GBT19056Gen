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
