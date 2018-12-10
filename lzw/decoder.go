package lzw

import (
	"../codesIO"
	"../header"
)

func decode(cr codesIO.Reader) []byte {
	dict := createDictionary()
	result := make([]byte, 0)
	buf := make([]byte, 0)
	for cr.HasCodes() {
		i := cr.Get(dict.getCodeSize() + codesIO.CodeHeadLength)
		
		// TODO: check code head
		
		i >>= codesIO.CodeHeadLength
		
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

func Decode(src *[]byte) (*[]byte, error) {
	if len(*src) == 0 {
		return src, nil
	}
	header.Fill(src)
	content := (*src)[header.GetLength():]
	if err := header.CheckPackedContent(&content); err != nil {
		return nil, err
	}
	res := decode(codesIO.AcquireCodes(&content))
	if err := header.CheckUnpackedContent(&res); err != nil {
		return nil, err
	}
	return &res, nil
}
