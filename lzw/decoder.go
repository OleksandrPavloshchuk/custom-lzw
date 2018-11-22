package lzw

import (
    "../codesIO"
)

func decode(cr codesIO.CodeReader) []byte {
	dict := createDictionary()
	result := make([]byte, 0)
	buf := make([]byte, 0)
	for cr.HasCodes() {
		i := cr.Get(dict.getCodeSize())
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

func Decode(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	h := header{&src}
	if err := h.CheckPackedContent(); err != nil {
		return nil, err
	}
	res := decode(codesIO.AcquireCodes(src, HeadLen))
	if err := h.CheckUnpackedContent(&res); err != nil {
		return nil, err
	}
	return res, nil
}
