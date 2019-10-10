package gbt19056

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

// DecodeGBK convert GBK to UTF-8
func DecodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// DecodeGBKStr convert GBK to UTF-8 String
func DecodeGBKStr(s []byte) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	length := len(s)
	// removing tailing zeros
	for i := length - 1; i >= 0; i-- {
		if s[i] != 0x00 {
			length = i + 1
			break
		}
		if i == 0 {
			return "", nil
		}
	}
	d, e := DecodeGBK(s[:length])
	if e != nil {
		return "", e
	}

	str := strings.TrimSpace(string(d))
	return str, nil
}

func bytesToStr(s []byte) string {
	if len(s) == 0 {
		return ""
	}
	length := len(s)
	// removing tailing zeros
	for i := length - 1; i >= 0; i-- {
		if s[i] != 0x00 {
			length = i + 1
			break
		}
		if i == 0 {
			return ""
		}
	}
	return string(s[:length])
}

// EncodeGBK convert UTF-8 to GBK
func EncodeGBK(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

//DecodeHZGB2312 convert GBK to UTF-8
func DecodeHZGB2312(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.HZGB2312.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// EncodeHZGB2312 convert UTF-8 to GBK
func EncodeHZGB2312(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, simplifiedchinese.HZGB2312.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// DecodeBig5 convert BIG5 to UTF-8
func DecodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewDecoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// EncodeBig5 convert UTF-8 to BIG5
func EncodeBig5(s []byte) ([]byte, error) {
	I := bytes.NewReader(s)
	O := transform.NewReader(I, traditionalchinese.Big5.NewEncoder())
	d, e := ioutil.ReadAll(O)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func float64ToBytes(v float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, v)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func int32ToBytes(v int32) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, v)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func int16ToBytes(v int16) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, v)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}

func bytesToInt32(b []byte) int32 {
	buf := bytes.NewBuffer(b)
	var x int32
	binary.Read(buf, binary.BigEndian, &x)
	return x
}

func bytesToInt16(b []byte) int16 {
	buf := bytes.NewBuffer(b)
	var x int16
	binary.Read(buf, binary.BigEndian, &x)
	return x
}
