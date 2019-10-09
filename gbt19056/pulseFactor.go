package gbt19056

// PulseFactor ..
type PulseFactor struct {
	dataBlockMeta
	Now    DateTime `json:"now,string"`
	Factor uint16   `json:"Factor"`
}
