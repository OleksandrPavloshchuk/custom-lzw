package lzw

import (
	"io/ioutil"
)

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

func Encode(inputFileName string, outputFileName string, version []byte) error {
	src, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		return err
	}
	codeWriter := CodeWriter{}
	encode(src, &codeWriter)
	res := codeWriter.GetBytes()
	setHeader(&src, &res, version)
	return ioutil.WriteFile(outputFileName, res, 0644)
}

func setHeader(src *[]byte, res *[]byte, version []byte) {
    h := Header{res}
	h.SetSignature()
	h.SetVersion(version)
	h.SetUnpackedSize(uint64(len(*src)))
	h.SetPackedSize()
	h.SetUnpackedCRC(src)
	h.SetPackedCRC()
}
