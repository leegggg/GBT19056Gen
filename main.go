// You can edit this code!
// Click here and start typing.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/leegggg/GBT19056Gen/gbt19056"
)

func jsonToBinary() {
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

	var nbWrite int
	f, err := os.Create("data/example.dat")
	defer f.Close()

	writer := bufio.NewWriter(f)
	nbWrite, err = writer.Write(bs)
	fmt.Printf("wrote %d bytes\n", nbWrite)
	writer.Flush()

	return
}

func binaryToJSON() {

	file, err := os.Open("data/example.dat")
	//file, err := os.Open("data/example.dat")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return
	}

	// calculate the bytes size
	var size int64 = info.Size()
	buff := make([]byte, size)

	// read into buffer
	reader := bufio.NewReader(file)
	_, err = reader.Read(buff)

	exportRecord := new(gbt19056.ExportRecord)
	exportRecord.LoadBinary(buff)
	fmt.Println(exportRecord.RealTime.Now)

	exportRecordJSON, _ := json.MarshalIndent(exportRecord, "", "    ")
	err = ioutil.WriteFile("data/output.json", exportRecordJSON, 0644)

	return
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Panic with %s, pls check input file or read help message.", err)
			os.Exit(-1)
		}
	}()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "%s [-o output] [-ts] <input>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "GB/T-19056 binary dump file r/w outil by ylin\n")
		flag.PrintDefaults()
	}
	hasTsPtr := flag.Bool("ts", false, "add a timestamp in the output file name(optional)")
	useUI := flag.Bool("i", false, "add a timestamp in the output file name(optional)")
	var output string
	flag.StringVar(&output, "o", "", "output file path(optional)")

	flag.Parse()
	if *useUI {
		transUI()
		return
	}

	other := flag.Args()

	if len(other) <= 0 {
		fmt.Fprintf(os.Stderr, "Need at least a input file path in cli\n")
		fmt.Fprintf(os.Stderr, "Try '%s --help' for more information.\n", os.Args[0])
		os.Exit(-1)
	}

	inputPath := other[0]

	// Open our jsonFile
	inFile, err := os.Open(inputPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Error when open file %s\n", inputPath))
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(-1)
	}
	fmt.Println(fmt.Sprintf("Successfully Opened %s\n", inputPath))
	// defer the closing of our jsonFile so that we can parse it later on
	defer inFile.Close()

	var bom [3]byte
	_, err = io.ReadFull(inFile, bom[:])
	if err != nil {
		log.Fatal(err)
	}
	if bom[0] == 0xef && bom[1] == 0xbb && bom[2] == 0xbf {
		fmt.Println("Warning: Json file with BOM detechted we can handle it but this is not recommended for json.")
	} else {
		_, err = inFile.Seek(0, 0) // Not a BOM -- seek back to the beginning
		if err != nil {
			log.Fatal(err)
		}
	}

	buff, _ := ioutil.ReadAll(inFile)

	exportRecord := new(gbt19056.ExportRecord)
	if filepath.Ext(inputPath) == ".json" {
		err = json.Unmarshal(buff, exportRecord)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error unmarshal json file\n")
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		fmt.Println("Unmarshal JSON Done")
	} else {
		err = exportRecord.LoadBinary(buff)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error decode gbt19056 binary file\n")
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		fmt.Println("Decode gbt19056 binary Done")
	}
	fmt.Println("Read Record Done")

	// Output file
	outputPath := exportRecord.MakeFileName()

	if *hasTsPtr {
		outputPath = fmt.Sprintf("%s.%s", outputPath, time.Now().Format("2006-01-02T150405"))
	}
	if filepath.Ext(inputPath) == ".json" {
		var bs []byte
		bs, err = exportRecord.DumpData()

		if len(output) > 0 {
			outputPath = output
		} else {
			outputPath = fmt.Sprintf("%s%s", outputPath, ".VDR")
		}
		// outputPath, err = filepath.Abs(outputPath)
		var nbWrite int
		nbWrite = len(bs)
		//var outputFile *os.File
		// outputFile, err = os.Create(outputPath)
		// defer outputFile.Close()

		// writer := bufio.NewWriter(outputFile)
		// nbWrite, err = writer.Write(bs)
		// err = writer.Flush()
		err = ioutil.WriteFile(outputPath, bs, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		fmt.Printf("Wrote %d bytes to binary %s\n", nbWrite, outputPath)

	} else {
		if len(output) > 0 {
			outputPath = output
		} else {
			outputPath = fmt.Sprintf("%s%s", outputPath, ".json")
		}
		// outputPath, err = filepath.Abs(outputPath)
		exportRecordJSON, _ := json.MarshalIndent(exportRecord, "", "    ")
		err = ioutil.WriteFile(outputPath, exportRecordJSON, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		fmt.Printf("Wrote JSON to  %s\n", outputPath)
	}
}
