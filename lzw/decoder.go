package lzw

import (
	"errors"
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

func Decode(src []byte, version []byte) ([]byte,error) {
	h := Header{&src}	
	if err := checkHeader(&h, version); err != nil {
		return nil, err
	}
	codeReader := CodeReader{}
	codeReader.Set(src)
	res:=decode(&codeReader)
	if err := checkUnpackedContent(&h, &res); err!=nil {
	    return nil, err
	}
	return res, nil
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
	if !h.CheckPackedCRC() {
	    return errors.New("invalid packed CRC")
	}
	return nil
}

func checkUnpackedContent(h *Header, res *[]byte) error {
	if !h.CheckUnpackedSize(uint64(len(*res))) {
	    return errors.New("invalid unpacked content size")
	}
	if !h.CheckUnpackedCRC(res) {
	    return errors.New("invalid unpacked content CRC")
	}
	return nil
}
