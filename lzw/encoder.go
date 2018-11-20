package lzw

func encode(src []byte, cw *codeWriter) {
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
				cw.accept(IncrementCodeSize, dict.getCodeSize()-1)
			}
			buf = make([]byte, 0)
		}
		buf = append(buf, b)
	}
	emit(buf, dict, cw)
}

func emit(s []byte, dict dictionary, cw *codeWriter) {
	cw.accept(dict.getIndex(s), dict.getCodeSize())
}

func Encode(src []byte) ([]byte, error) {
	cw := codeWriter{}
	encode(src, &cw)
	res := cw.getBytes()
	if len(res) == HeadLen {
		return []byte{}, nil
	}
	setHeader(&res, &src)
	return res, nil
}
