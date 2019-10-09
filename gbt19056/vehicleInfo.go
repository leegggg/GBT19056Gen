package gbt19056

// VehicleInfo ..
type VehicleInfo struct {
	dataBlockMeta
	ID        string `json:"id"`
	Plate     string `json:"plate"`
	PlateType string `json:"plate_type"`
}
