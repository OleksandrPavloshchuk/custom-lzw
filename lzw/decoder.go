package lzw

import (
	"errors"
	"io/ioutil"
)

var VersionChecker func(int,*[]byte) bool

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

func Decode(inputFileName string, outputFileName string, version []byte) error {
	src, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		return err
	}
	h := Header{&src}
	
	if err := checkHeader(&h, version); err != nil {
		return err
	}
	codeReader := CodeReader{}
	codeReader.Set(src)
	res:=decode(&codeReader)
	if !h.CheckUnpackedSize(uint64(len(res))) {
	    return errors.New("invalid unpacked content size")
	}
	// - TODO CRC
		
	return ioutil.WriteFile(outputFileName, res, 0644)
}

func checkHeader(h *Header, version []byte) error {
	if !h.CheckSignature() {
		return errors.New("invalid archive signature")
	}
	if !h.CheckVersion(VersionChecker) {
		return errors.New("invalid archive version")
	}
	if !h.CheckPackedSize() {
	    return errors.New("invalid packed content size")
	}
	return nil

}
