package gbt19056

import "github.com/leegggg/GBT19056Gen/utils/bcd"

// OdometerMultiplier ...
const OdometerMultiplier = 10

// Odometer ..
type Odometer struct {
	dataBlockMeta
	Now          DateTime `json:"now,string"`
	TimeInit     DateTime `json:"time_init,string"`
	MileageTotal float64  `json:"mileage_total"`
	MileageInit  float64  `json:"mileage_init"`
}

// DumpData Odometer
func (e *Odometer) DumpData() ([]byte, error) {
	var now, timeInit, mileageTotal, mileageInit []byte
	var err error

	now, err = e.Now.DumpData()
	timeInit, err = e.TimeInit.DumpData()
	mileageInit = bcd.FromUint32(uint32(e.MileageInit * OdometerMultiplier))
	mileageTotal = bcd.FromUint32(uint32(e.MileageTotal * OdometerMultiplier))

	bs := append(now, timeInit...)
	bs = append(bs, mileageInit...)
	bs = append(bs, mileageTotal...)
	bs, err = e.linkDumpedData(bs)
	return bs, err
}

// LoadBinary RealTime Table A.9, Code 0x03
func (e *Odometer) LoadBinary(buffer []byte, meta dataBlockMeta) {
	e.dataBlockMeta = meta
	e.Now.LoadBinary(buffer[0:6])
	e.TimeInit.LoadBinary(buffer[6:12])
	e.MileageInit = float64(bcd.ToUint32(buffer[12:16])) / OdometerMultiplier
	e.MileageTotal = float64(bcd.ToUint32(buffer[16:20])) / OdometerMultiplier
	return
}
