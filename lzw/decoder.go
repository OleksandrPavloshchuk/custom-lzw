package lzw

var VersionChecker func(int,*[]byte) bool

func decode(cr *codeReader) []byte {
	dict := createDictionary()
	result := make([]byte, 0)
	buf := make([]byte, 0)
	for cr.hasCodes() {
		i := cr.get(dict.getCodeSize())
		if !dict.incrementCodeSizeWhileDecode(i) {
			var s []byte
			if dict.hasCode(i) {
				s = dict.getString(i)
			} else {
				s = append(buf, buf[0])
			}
			test := append(buf, s...)
			if !dict.hasString(test) {
				result = append(result, buf...)
				buf = append(buf, s[0])
				dict.put(buf)
				buf = make([]byte, 0)
			}
			buf = append(buf, s...)
		}
	}
	result = append(result, buf...)
	return result
}

func Decode(src []byte, version []byte) ([]byte,error) {
    if len(src)==0 {
        return []byte{},nil
    }
	h := header{&src}	
	if err := h.CheckPackedContent(version); err != nil {
		return nil, err
	}
	cr := codeReader{}
	cr.set(src)
	res:=decode(&cr)
	if err := h.CheckUnpackedContent(&res); err!=nil {
	    return nil, err
	}
	return res, nil
}
