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
