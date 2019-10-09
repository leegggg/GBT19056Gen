package gbt19056

// SpeedStatusLog ..
type SpeedStatusLog struct {
	dataBlockMeta
	Records []SpeedStatusLogRecord `json:"records"`
}

// SpeedStatusLogRecord ...
type SpeedStatusLogRecord struct {
	Status HexUint8 `json:"status"` // 0x01: normal; 0x02: error;
	Start  DateTime `json:"start,string"`
	End    DateTime `json:"end,string"`
	Speeds []uint8  `json:"speeds"`
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
