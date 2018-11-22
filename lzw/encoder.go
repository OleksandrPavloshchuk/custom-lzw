package lzw

import (
    "../codesIO"
)

func encode(src []byte, cw *codesIO.CodeWriter) {
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

func emit(s []byte, dict dictionary, cw *codesIO.CodeWriter) {
	cw.Accept(dict.getIndex(s), dict.getCodeSize())
}

func Encode(src []byte) ([]byte, error) {
	cw := codesIO.CodeWriter{}
	encode(src, &cw)
	res := cw.GetBytes(HeadLen)
	if len(res) == HeadLen {
		return []byte{}, nil
	}
	setHeader(&res, &src)
	return res, nil
}
