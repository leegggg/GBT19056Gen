package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Checkin struct {
	Timestamp SpecialDate `json:"timestamp"`
	User      string      `json:"user"`
}

type SpecialDate struct {
	time.Time
}

func (sd *SpecialDate) UnmarshalJSON(input []byte) error {
	strInput := string(input)
	strInput = strings.Trim(strInput, `"`)
	newTime, err := time.Parse("2006/01/02 15:04:05", strInput)
	if err != nil {
		return err
	}

	sd.Time = newTime
	return nil
}

func main() {
	j := `{"timestamp":"2016/11/02 08:18:20", "user":"John Doe"}`

	var c Checkin
	err := json.Unmarshal([]byte(j), &c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c)
}
