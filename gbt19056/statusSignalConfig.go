package gbt19056

// StatusSignalConfig ..
type StatusSignalConfig struct {
	dataBlockMeta
	Now    DateTime `json:"now,string"`
	Status Status   `json:"status"`
	Config []string `json:"config"`
}
