package gbt19056

// Odometer ..
type Odometer struct {
	dataBlockMeta
	Now          DateTime `json:"now,string"`
	TimeInit     DateTime `json:"time_init,string"`
	MileageTotal uint32   `json:"mileage_total"`
	MileageInit  uint32   `json:"mileage_init"`
}
