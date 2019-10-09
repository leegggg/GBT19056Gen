package gbt19056

// RecoderID ..
type RecoderID struct {
	dataBlockMeta
	CCC     string   `json:"CCC"`
	Version string   `json:"version"`
	Dop     DateTime `json:"dop,string"`
	Sn      uint32   `json:"sn"`
	Comment string   `json:"comment"`
}
