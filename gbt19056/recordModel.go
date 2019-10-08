package gbt19056

import (
	"encoding/binary"
	"strconv"
	"strings"
	"time"
)

// HexUint8 ...
type HexUint8 uint8

// Status ...
type Status [8]bool

// SpeedStatus ...
type SpeedStatus struct {
	Speed  uint8  `json:"speed"`
	Status Status `json:"status"`
}

// SpeedLogRecord ...
type SpeedLogRecord struct {
	Start         time.Time     `json:"start,string"`
	SpeedStatuses []SpeedStatus `json:"speed_statuses"`
}

// PositionLogRecord ...
type PositionLogRecord struct {
	Start     string `json:"start"`
	Positions []struct {
		Speed    uint8    `json:"speed"`
		Position Position `json:"position"`
	} `json:"positions"`
}

// Position ...
type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Elevation float64 `json:"elevation"`
}

// AccidentLogRecord ...
type AccidentLogRecord struct {
	Ts            time.Time     `json:"ts,string"`
	EndPosition   Position      `json:"position"`
	SpeedStatuses []SpeedStatus `json:"speed_statuses"`
}

// OvertimeLogRecord ...
type OvertimeLogRecord struct {
	Start         time.Time `json:"start,string"`
	End           time.Time `json:"end,string"`
	PositionStart Position  `json:"position_start"`
	PositionEnd   Position  `json:"position_end"`
}

// DriverLogRecord ...
type DriverLogRecord struct {
	Ts   time.Time `json:"ts,string"`
	ID   string    `json:"id"`
	Type HexUint8  `json:"type"` // 0x01: Login; 0x02: Logout; Other resversed
}

// ExternalPowerLogRecord ...
type ExternalPowerLogRecord struct {
	Ts   time.Time `json:"ts,string"`
	Type HexUint8  `json:"type"` // 0x01: pluged; 0x02: unpluged;
}

// ConfigChangeLogRecord ...
type ConfigChangeLogRecord struct {
	Ts   time.Time `json:"ts,string"`
	Type HexUint8  `json:"event"`
}

// SpeedStatusLogRecord ...
type SpeedStatusLogRecord struct {
	Status HexUint8  `json:"status"` // 0x01: normal; 0x02: error;
	Start  time.Time `json:"start,string"`
	End    time.Time `json:"end,string"`
	Speeds []uint8   `json:"speeds"`
}

// dataBlockMeta ...
type dataBlockMeta struct {
	Code HexUint8 `json:"code"` // need decode from 0x01 ...
	Name string   `json:"name"`
}

// DumpDate ...
func (e *dataBlockMeta) DumpDate() ([]byte, error) {
	bs := make([]byte, 19)
	bs[0] = (byte)(e.Code)
	sub := bs[1:]
	name, _ := EncodeGBK(([]byte)(e.Name))
	copy(sub, name)
	return bs, nil
}

// DriverInfo ..
type DriverInfo struct {
	dataBlockMeta
	ID string `json:"id"`
}

// RealTime ..
type RealTime struct {
	dataBlockMeta
	Now time.Time `json:"now,string"`
}

// Odometer ..
type Odometer struct {
	dataBlockMeta
	Now          time.Time `json:"now,string"`
	TimeInit     time.Time `json:"time_init,string"`
	MileageTotal uint32    `json:"mileage_total"`
	MileageInit  uint32    `json:"mileage_init"`
}

// PulseFactor ..
type PulseFactor struct {
	dataBlockMeta
	Now    time.Time `json:"now,string"`
	Factor uint16    `json:"Factor"`
}

// VehicleInfo ..
type VehicleInfo struct {
	dataBlockMeta
	ID        string `json:"id"`
	Plate     string `json:"plate"`
	PlateType string `json:"plate_type"`
}

// StatusSignalConfig ..
type StatusSignalConfig struct {
	dataBlockMeta
	Now    time.Time `json:"now,string"`
	Status Status    `json:"status"`
	Config []string  `json:"config"`
}

// RecoderID ..
type RecoderID struct {
	dataBlockMeta
	CCC     string    `json:"CCC"`
	Version string    `json:"version"`
	Dop     time.Time `json:"dop,string"`
	Sn      uint32    `json:"sn"`
	Comment string    `json:"comment"`
}

// SpeedLog ..
type SpeedLog struct {
	dataBlockMeta
	Records []SpeedLogRecord `json:"records"`
}

// PositionLog ..
type PositionLog struct {
	dataBlockMeta
	Records []PositionLogRecord `json:"records"`
}

// AccidentLog ..
type AccidentLog struct {
	dataBlockMeta
	Records []AccidentLogRecord `json:"records"`
}

// OvertimeLog ..
type OvertimeLog struct {
	dataBlockMeta
	Records []OvertimeLogRecord `json:"records"`
}

// DriverLog ..
type DriverLog struct {
	dataBlockMeta
	Records []DriverLogRecord `json:"records"`
}

// ExternalPowerLog ..
type ExternalPowerLog struct {
	dataBlockMeta
	Records []ExternalPowerLogRecord `json:"records"`
}

// ConfigChangeLog ..
type ConfigChangeLog struct {
	dataBlockMeta
	Records []ConfigChangeLogRecord `json:"records"`
}

// SpeedStatusLog ..
type SpeedStatusLog struct {
	dataBlockMeta
	Records []SpeedStatusLogRecord `json:"records"`
}

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

// DumpDate ...
func (e *ExportRecord) DumpDate() ([]byte, error) {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, e.NumberBlock)
	standardVersionDump, _ := e.StandardVersion.DumpDate()
	bs = append(bs, standardVersionDump...)
	return bs, nil
}

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

// UnmarshalJSON ...
// func (s *SpeedStatusLog) UnmarshalJSON(data []byte) error {
// 	type Alias SpeedStatusLog
// 	aux := &struct {
// 		*Alias
// 	}{
// 		Alias: (*Alias)(s),
// 	}
// 	if err := json.Unmarshal(data, &aux); err != nil {
// 		return err
// 	}
// 	// s = SpeedStatusLog(*aux)
// 	// s.SomeCustomType = time.Unix(aux.SomeCustomType, 0)
// 	return nil
// }
// 	return nil
// }
// }
// 	// s.SomeCustomType = time.Unix(aux.SomeCustomType, 0)
// 	return nil
// }
