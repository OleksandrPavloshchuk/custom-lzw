package lzw

import (
	"errors"
	"io/ioutil"
)

func decode(codeReader *CodeReader) []byte {
	dict := Dictionary{}
	dict.Init()

	result := make([]byte, 0)
	buf := make([]byte, 0)
	for codeReader.HasCodes() {
		i := codeReader.Get(dict.GetCodeSize())
		if !dict.IncrementCodeSizeWhileDecode(i) {
			var s []byte
			if dict.HasCode(i) {
				s = dict.GetString(i)
			} else {
				s = append(buf, buf[0])
			}
			test := append(buf, s...)
			if !dict.HasString(test) {
				result = append(result, buf...)
				buf = append(buf, s[0])
				dict.Put(buf)
				buf = make([]byte, 0)
			}
			buf = append(buf, s...)
		}
	}
	result = append(result, buf...)
	return result
}

func Decode(inputFileName string, outputFileName string) error {
	src, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		return err
	}
	if err := checkHeader(&src); err != nil {
		return err
	}
	codeReader := CodeReader{}
	codeReader.Set(src)
	return ioutil.WriteFile(outputFileName, decode(&codeReader), 0644)
}

func checkHeader(src *[]byte) error {
	if !CheckSignature(src) {
		return errors.New("invalid archive signature")
	}
	// - TODO version
	// - TODO unpacked size
	// - TODO packed size
	// - TODO CRC
	return nil

}
