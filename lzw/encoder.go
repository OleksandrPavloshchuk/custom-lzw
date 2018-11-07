package lzw

func encode(src []byte, codeWriter *CodeWriter) {
	dict := Dictionary{}
	dict.Init()

	buf := make([]byte, 0)
	for _, b := range src {
		test := append(buf, b)
		if !dict.HasString(test) {
			emit(buf, dict, codeWriter)
			dict.Put(test)
			if dict.IncrementCodeSizeWhileEncode() {
				codeWriter.Accept(IncrementCodeSize, dict.GetCodeSize()-1)
			}
			buf = make([]byte, 0)
		}
		buf = append(buf, b)
	}
	emit(buf, dict, codeWriter)
}

func emit(s []byte, dict Dictionary, codeWriter *CodeWriter) {
	codeWriter.Accept(dict.GetIndex(s), dict.GetCodeSize())
}

func Encode(src []byte, version []byte) ([]byte, error) {
	codeWriter := CodeWriter{}
	encode(src, &codeWriter)
	res := codeWriter.GetBytes()
	SetHeader(&res, &src, version)
	return res, nil
}

