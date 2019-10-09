package gbt19056

// VehicleInfo ..
type VehicleInfo struct {
	dataBlockMeta
	ID        string `json:"id"`
	Plate     string `json:"plate"`
	PlateType string `json:"plate_type"`
}

// DumpData VehicleInfo
func (e *VehicleInfo) DumpData() ([]byte, error) {
	var buff []byte
	var err error

	id := make([]byte, 17)
	// ASCII ID should be safe to be copy directly
	copy(id, []byte(e.ID))

	// TODO: Check length
	plate := make([]byte, 12)
	buff, err = EncodeGBK(([]byte)(e.Plate))
	copy(plate, buff)

	plateType := make([]byte, 12)
	buff, err = EncodeGBK(([]byte)(e.PlateType))
	copy(plateType, buff)

	bs := append(id, plate...)
	bs = append(bs, plateType...)

	bs, err = e.linkDumpedData(bs)
	return bs, err
}
