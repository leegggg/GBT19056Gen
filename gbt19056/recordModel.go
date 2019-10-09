package gbt19056

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/leegggg/GBT19056Gen/utils/bcd"
)

// HexUint8 ...
type HexUint8 uint8

// UnmarshalJSON HexUint8 ...
func (sd *HexUint8) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	res, err := strconv.ParseUint(strInput, 0, 8)
	if err != nil {
		return err
	}

	*sd = HexUint8(res)
	return nil
}

// DateTime ...
type DateTime struct {
	time.Time
}

// DumpData DateTime
func (e *DateTime) DumpData() ([]byte, error) {
	bs := make([]byte, 6)
	if e.Year() < 2000 || e.Year() > 2099 {
		err := fmt.Errorf("Year %d not in range [2000-2099]", e.Year())
		return bs, err
	}

	var nb uint8
	nb = uint8(e.Year() - 2000)
	bs[0] = bcd.FromUint8(nb)

	nb = uint8(e.Month())
	bs[1] = bcd.FromUint8(nb)

	nb = uint8(e.Day())
	bs[2] = bcd.FromUint8(nb)

	nb = uint8(e.Hour())
	bs[3] = bcd.FromUint8(nb)

	nb = uint8(e.Minute())
	bs[4] = bcd.FromUint8(nb)

	nb = uint8(e.Second())
	bs[5] = bcd.FromUint8(nb)
	return bs, nil
}

// DumpDataShort DateTime
func (e *DateTime) DumpDataShort() ([]byte, error) {
	bs := make([]byte, 3)
	if e.Year() < 2000 || e.Year() > 2099 {
		err := fmt.Errorf("Year %d not in range [2000-2099]", e.Year())
		return bs, err
	}

	var nb uint8
	nb = uint8(e.Year() - 2000)
	bs[0] = bcd.FromUint8(nb)

	nb = uint8(e.Month())
	bs[1] = bcd.FromUint8(nb)

	nb = uint8(e.Day())
	bs[2] = bcd.FromUint8(nb)

	return bs, nil
}

// Status ...
type Status [8]bool

// DumpData DateTime
func (e *Status) DumpData() (byte, error) {
	var bs byte
	bs = 0x00
	if len(*e) > 8 {
		return bs, fmt.Errorf("Status array too long")
	}
	// Set bitmap
	for index, bit := range *e {
		if bit {
			bs |= 1 << index
		}
	}
	return bs, nil
}

// SpeedStatus ...
type SpeedStatus struct {
	Speed  uint8  `json:"speed"`
	Status Status `json:"status"`
}

// DumpData SpeedStatus
func (e *SpeedStatus) DumpData() ([]byte, error) {
	bs := make([]byte, 2)
	bs[0] = e.Speed
	status, _ := e.Status.DumpData()
	bs[1] = status
	return bs, nil
}

// Position ...
type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Elevation float64 `json:"elevation"`
}

// DumpData Position
func (e *Position) DumpData() ([]byte, error) {
	bs := make([]byte, 10)
	positionMultiplier := 10000.0
	latitude := int32(math.Round(e.Latitude * positionMultiplier))
	longitude := int32(math.Round(e.Longitude * positionMultiplier))
	elevation := int16(e.Elevation)

	copy(bs[0:4], int32ToBytes(latitude))
	copy(bs[4:8], int32ToBytes(longitude))
	copy(bs[8:], int16ToBytes(elevation))
	return bs, nil
}

// dataBlockMeta ...
type dataBlockMeta struct {
	Code HexUint8 `json:"code"` // need decode from 0x01 ...
	Name string   `json:"name"`
}

// DumpDate ...
func (e *dataBlockMeta) DumpData() ([]byte, error) {
	bs := make([]byte, 19)
	bs[0] = (byte)(e.Code)
	sub := bs[1:]
	name, _ := EncodeGBK(([]byte)(e.Name))
	copy(sub, name)
	return bs, nil
}

// linkDumpedData
func (e dataBlockMeta) linkDumpedData(body []byte) ([]byte, error) {
	meta, err := e.DumpData()
	// Data Length
	dataLength := make([]byte, 4)
	binary.BigEndian.PutUint32(dataLength, (uint32)(len(body)))
	meta = append(meta, dataLength...)
	bs := append(meta, body...)
	return bs, err
}
