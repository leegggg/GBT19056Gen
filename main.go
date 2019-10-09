// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/leegggg/GBT19056Gen/gbt19056"
)

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("data/example.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	jsonBlob, _ := ioutil.ReadAll(jsonFile)

	var record gbt19056.ExportRecord
	err = json.Unmarshal(jsonBlob, &record)

	fmt.Printf("0x%02x %s", record.SpeedStatusLog.Code, record.SpeedStatusLog.Name)
	fmt.Println(record.RealTime.Now)

	var bs []byte
	bs, err = record.DumpData()
	print(bs)

	var nbWrite int
	f, err := os.Create("data/example.dat")
	defer f.Close()

	writer := bufio.NewWriter(f)
	nbWrite, err = writer.Write(bs)
	fmt.Printf("wrote %d bytes\n", nbWrite)
	writer.Flush()

	return
}
