package lzw

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
	if err := h.CheckPackedContent(version); err != nil {
		return nil, err
	}
	codeReader := CodeReader{}
	codeReader.Set(src)
	res:=decode(&codeReader)
	if err := h.CheckUnpackedContent(&res); err!=nil {
	    return nil, err
	}
	return res, nil
}
