package gbt19056

import "encoding/binary"

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
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, e.NumberBlock)
	standardVersionDump, _ := e.StandardVersion.DumpData()
	bs = append(bs, standardVersionDump...)
	return bs, nil
}
