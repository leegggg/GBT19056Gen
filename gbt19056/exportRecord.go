package gbt19056

import (
	"encoding/binary"
	"fmt"
)

// ExportRecord ...
type ExportRecord struct {
	NumberBlock        uint16             `json:"number_block"`
	StandardVersion    StandardVersion    `json:"standard_version"`
	DriverInfo         DriverInfo         `json:"driver_info"`
	RealTime           RealTime           `json:"real_time"`
	Odometer           Odometer           `json:"odometer"`
	PulseFactor        PulseFactor        `json:"pulse_factor"`
	VehicleInfo        VehicleInfo        `json:"vehicle_info"`
	StatusSignalConfig StatusSignalConfig `json:"status_signal_config"`
	RecoderID          RecoderID          `json:"recoder_id"`
	SpeedLog           SpeedLog           `json:"speed_log"`
	PositionLog        PositionLog        `json:"position_log"`
	AccidentLog        AccidentLog        `json:"accident_log"`
	OvertimeLog        OvertimeLog        `json:"overtime_log"`
	DriverLog          DriverLog          `json:"driver_log"`
	ExternalPowerLog   ExternalPowerLog   `json:"external_power_log"`
	ConfigChangeLog    ConfigChangeLog    `json:"config_change_log"`
	SpeedStatusLog     SpeedStatusLog     `json:"speed_status_log"`
}

// DumpData ExportRecord
func (e *ExportRecord) DumpData() ([]byte, error) {
	var buff []byte
	var err error

	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, e.NumberBlock)

	buff, err = e.StandardVersion.DumpData()
	bs = append(bs, buff...)
	buff, err = e.DriverInfo.DumpData()
	bs = append(bs, buff...)
	buff, err = e.RealTime.DumpData()
	bs = append(bs, buff...)
	buff, err = e.Odometer.DumpData()
	bs = append(bs, buff...)
	buff, err = e.PulseFactor.DumpData()
	bs = append(bs, buff...)
	buff, err = e.VehicleInfo.DumpData()
	bs = append(bs, buff...)
	buff, err = e.StatusSignalConfig.DumpData()
	bs = append(bs, buff...)
	buff, err = e.RecoderID.DumpData()
	bs = append(bs, buff...)
	buff, err = e.SpeedLog.DumpData()
	bs = append(bs, buff...)
	buff, err = e.PositionLog.DumpData()
	bs = append(bs, buff...)
	buff, err = e.AccidentLog.DumpData()
	bs = append(bs, buff...)
	buff, err = e.OvertimeLog.DumpData()
	bs = append(bs, buff...)
	buff, err = e.DriverLog.DumpData()
	bs = append(bs, buff...)
	buff, err = e.ExternalPowerLog.DumpData()
	bs = append(bs, buff...)
	buff, err = e.ConfigChangeLog.DumpData()
	bs = append(bs, buff...)
	buff, err = e.SpeedStatusLog.DumpData()
	bs = append(bs, buff...)

	var checkSum uint8
	checkSum = 0x00

	for _, v := range bs {
		checkSum ^= v
	}

	bs = append(bs, checkSum)

	return bs, err
}

// LoadBinary ExportRecord
func (e *ExportRecord) LoadBinary(buffer []byte) {
	var offset int
	var err error
	var data []byte
	ptr := 0
	e.NumberBlock = binary.BigEndian.Uint16(buffer[ptr : ptr+2])
	ptr += 2
	for indexBlock := uint16(0); indexBlock < e.NumberBlock; indexBlock++ {
		meta := new(dataBlockMeta)
		offset, err = meta.LoadBinary(buffer[ptr:])
		data = buffer[ptr+LengthMetadata : ptr+offset]
		switch code := meta.Code; code {
		case HexUint8(0x00):
			e.StandardVersion.LoadBinary(data, *meta)
		case HexUint8(0x01):
			e.DriverInfo.LoadBinary(data, *meta)
		case HexUint8(0x02):
			e.RealTime.LoadBinary(data, *meta)
		case HexUint8(0x03):
			e.Odometer.LoadBinary(data, *meta)
		case HexUint8(0x04):
			e.PulseFactor.LoadBinary(data, *meta)
		case HexUint8(0x05):
			e.VehicleInfo.LoadBinary(data, *meta)
		case HexUint8(0x06):
			e.StatusSignalConfig.LoadBinary(data, *meta)
		case HexUint8(0x07):
			e.RecoderID.LoadBinary(data, *meta)
		case HexUint8(0x08):
			e.SpeedLog.LoadBinary(data, *meta)
		case HexUint8(0x09):
			e.PositionLog.LoadBinary(data, *meta)
		case HexUint8(0x10):
			e.AccidentLog.LoadBinary(data, *meta)
		case HexUint8(0x11):
			e.OvertimeLog.LoadBinary(data, *meta)
		case HexUint8(0x12):
			e.DriverLog.LoadBinary(data, *meta)
		case HexUint8(0x13):
			e.ExternalPowerLog.LoadBinary(data, *meta)
		case HexUint8(0x14):
			e.ConfigChangeLog.LoadBinary(data, *meta)
		case HexUint8(0x15):
			e.SpeedStatusLog.LoadBinary(data, *meta)
		default:
			// freebsd, openbsd,
			// plan9, windows...
			fmt.Printf("%d.\n", code)
		}
		ptr += offset
	}
	if err != nil {

	}
}
