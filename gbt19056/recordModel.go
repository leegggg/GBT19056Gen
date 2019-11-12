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

// LengthMetadata ...
const LengthMetadata = 23

// PositionMultiplier ...
const PositionMultiplier = 10000.0 * 60

// HexUint8 ...
type HexUint8 uint8

// UnmarshalJSON HexUint8 ...
func (sd *HexUint8) UnmarshalJSON(input []byte) error {
	strInput := bytesToStr(input)
	strInput = strings.Trim(strInput, `"`)
	res, err := strconv.ParseUint(strInput, 0, 8)
	if err != nil {
		return err
	}

	*sd = HexUint8(res)
	return nil
}

// MarshalJSON HexUint8
func (sd *HexUint8) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := "\"0x00\""
	if uint8(*sd) != 0x00 {
		stamp = fmt.Sprintf("\"0x%02x\"", uint8(*sd))
	}
	return []byte(stamp), nil
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

// LoadBinary RealTime Table A.8, Code 0x02
func (e *DateTime) LoadBinary(buffer []byte) {
	year := bcd.ToUint8(buffer[0])
	month := bcd.ToUint8(buffer[1])
	day := bcd.ToUint8(buffer[2])
	hour := bcd.ToUint8(buffer[3])
	min := bcd.ToUint8(buffer[4])
	sec := bcd.ToUint8(buffer[5])
	e.Time = time.Date(
		int(year)+2000, time.Month(int(month)), int(day), int(hour), int(min), int(sec), 0, time.UTC)
}

// LoadBinaryShort RealTime Table A.14
func (e *DateTime) LoadBinaryShort(buffer []byte) {
	year := bcd.ToUint8(buffer[0])
	month := bcd.ToUint8(buffer[1])
	day := bcd.ToUint8(buffer[2])

	e.Time = time.Date(
		int(year)+2000, time.Month(int(month)), int(day), 0, 0, 0, 0, time.UTC)
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

// LoadBinary Status Table A.8, Code 0x02
func (e *Status) LoadBinary(buffer byte) {
	for i := 0; i < 8; i++ {
		if (buffer & (1 << i)) == 0 {
			e[i] = false
		} else {
			e[i] = true
		}
	}
}

// SpeedStatus ...
type SpeedStatus struct {
	Speed  uint8  `json:"speed"`
	Status Status `json:"status"`
}

// LengthSpeedStatus ...
const LengthSpeedStatus = 2

// DumpData SpeedStatus
func (e *SpeedStatus) DumpData() ([]byte, error) {
	bs := make([]byte, 2)
	bs[0] = e.Speed
	status, _ := e.Status.DumpData()
	bs[1] = status
	return bs, nil
}

// LoadBinary SpeedStatus
func (e *SpeedStatus) LoadBinary(buffer []byte) {
	e.Speed = buffer[0]
	e.Status.LoadBinary(buffer[1])
}

// Position ...
type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Elevation float64 `json:"elevation"`
}

// LengthPosition ...
const LengthPosition = 10

// DumpData Position
func (e *Position) DumpData() ([]byte, error) {
	var longitude, latitude int32
	var elevation int16

	bs := make([]byte, 10)

	if e.Latitude == 0 && e.Longitude == 0 && e.Elevation == 0 {
		longitude = 0x7FFFFFFF
		latitude = 0x7FFFFFFF
		elevation = 0x7FFF
	} else {
		if math.Abs(e.Latitude) > 90 {
			latitude = 0x7FFFFFFF
		} else {
			latitude = int32(math.Round(e.Longitude * PositionMultiplier))
		}

		if math.Abs(e.Longitude) > 180 {
			longitude = 0x7FFFFFFF
		} else {
			longitude = int32(math.Round(e.Latitude * PositionMultiplier))
		}

		if math.Abs(e.Elevation) > 10000 {
			elevation = 0x7FFF
		} else {
			elevation = int16(e.Elevation)
		}

	}

	copy(bs[0:4], int32ToBytes(latitude))
	copy(bs[4:8], int32ToBytes(longitude))
	copy(bs[8:], int16ToBytes(elevation))
	return bs, nil
}

// LoadBinary ...
func (e *Position) LoadBinary(buffer []byte) {
	// Table A.20
	e.Longitude = float64(bytesToInt32(buffer[0:4])) / PositionMultiplier
	e.Latitude = float64(bytesToInt32(buffer[4:8])) / PositionMultiplier
	e.Elevation = float64(bytesToInt16(buffer[8:10]))
	return
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

// DumpDate ...
func (e *dataBlockMeta) LoadBinary(buffer []byte) (int, error) {
	// Table B.2
	var err error
	var name string
	dataLength := LengthMetadata
	e.Code = HexUint8(buffer[0])
	name, err = DecodeGBKStr(buffer[1:19])
	e.Name = name
	dataLength += int(binary.BigEndian.Uint32(buffer[19:23]))
	return dataLength, err
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
