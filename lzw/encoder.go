package lzw

import (
	"../codesIO"
	"../header"
)

func encode(src *[]byte, cw *codesIO.Writer) {
	if len(*src) == 0 {
		return
	}
	dict := createDictionary()
	buf := make([]byte, 0)
	for _, b := range *src {
		test := append(buf, b)
		if !dict.hasString(test) {
			emit(buf, dict, cw)
			dict.put(test)
			if dict.incrementCodeSizeWhileEncode() {
				cw.Accept(IncrementCodeSize, codesIO.CodeHeadLength)
			}
			buf = make([]byte, 0)
		}
		buf = append(buf, b)
	}
	emit(buf, dict, cw)
}

func emit(s []byte, dict dictionary, cw *codesIO.Writer) {
    code := dict.getIndex(s) 
    
    codeSize := dict.getCodeSize()
    codeHead := uint(2)
    if code<=255 {
        codeHead = 1
        codeSize = 8   
    }    
    code <<= codesIO.CodeHeadLength
    code += codeHead
    
    codeSize += codesIO.CodeHeadLength

	cw.Accept(code, codeSize)
}

func Encode(src *[]byte) (*[]byte, error) {
	cw := codesIO.Writer{}
	encode(src, &cw)
	res := cw.GetBytes()
	if len(res) == 0 {
		return &res, nil
	}
	header.SetSignature()
	header.SetVersion()
	header.SetUnpackedInfo(src)
	header.SetPackedInfo(&res)
	return &res, nil
}
