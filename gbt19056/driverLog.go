package gbt19056

// DriverLog ..
type DriverLog struct {
	dataBlockMeta
	Records []DriverLogRecord `json:"records"`
}

// DriverLogRecord ...
type DriverLogRecord struct {
	Ts   DateTime `json:"ts,string"`
	ID   string   `json:"id"`
	Type HexUint8 `json:"type"` // 0x01: Login; 0x02: Logout; Other resversed
}
