package gbt19056

import "github.com/leegggg/GBT19056Gen/utils/bcd"

// Odometer ..
type Odometer struct {
	dataBlockMeta
	Now          DateTime `json:"now,string"`
	TimeInit     DateTime `json:"time_init,string"`
	MileageTotal uint32   `json:"mileage_total"`
	MileageInit  uint32   `json:"mileage_init"`
}

// DumpData Odometer
func (e *Odometer) DumpData() ([]byte, error) {
	var now, timeInit, mileageTotal, mileageInit []byte
	var err error

	now, err = e.Now.DumpData()
	timeInit, err = e.TimeInit.DumpData()
	mileageInit = bcd.FromUint32(e.MileageInit)
	mileageTotal = bcd.FromUint32(e.MileageTotal)

	bs := append(now, timeInit...)
	bs = append(bs, mileageInit...)
	bs = append(bs, mileageTotal...)
	bs, err = e.linkDumpedData(bs)
	return bs, err
}
