package lzw

import (
    "../codesIO"
    "../header"
)

func encode(src []byte, cw *codesIO.Writer) {
	if len(src) == 0 {
		return
	}
	dict := createDictionary()
	buf := make([]byte, 0)
	for _, b := range src {
		test := append(buf, b)
		if !dict.hasString(test) {
			emit(buf, dict, cw)
			dict.put(test)
			if dict.incrementCodeSizeWhileEncode() {
				cw.Accept(IncrementCodeSize, dict.getCodeSize()-1)
			}
			buf = make([]byte, 0)
		}
		buf = append(buf, b)
	}
	emit(buf, dict, cw)
}

func emit(s []byte, dict dictionary, cw *codesIO.Writer) {
	cw.Accept(dict.getIndex(s), dict.getCodeSize())
}

func Encode(src []byte) ([]byte, error) {
	cw := codesIO.Writer{}
	encode(src, &cw)
	res := cw.GetBytes(0)
	if len(res) == 0 {
		return []byte{}, nil
	}
	header.SetSignature()
	header.SetVersion()
	header.SetUnpackedInfo(&src)
	return res, nil
}
