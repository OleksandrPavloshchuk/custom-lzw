package lzw

import (
	"io/ioutil"
)

type CodeWriter struct {
	CodeIO
}

func (this *CodeWriter) Accept(code uint, length uint) {
	d := uint(1)
	for i := uint(0); i < length; i++ {
		if code&d != 0 {
			this.bitSet.Set(this.start)
		}
		this.start++
		d <<= 1
	}
}

func (this *CodeWriter) Write(fileName string) error {
	size := (this.start + 7) / 8
	result := make([]byte, HeadLen)
	for i := 0; uint(i) < size; i++ {
		result = append(result, this.toByte(uint(i<<3)))
	}
	return ioutil.WriteFile(fileName, result, 0644)
}

func (this *CodeWriter) toByte(offset uint) byte {
	r := byte(0)
	d := byte(1)
	for i := 0; i < 8; i++ {
		if this.bitSet.IsSet(uint(i) + offset) {
			r |= d
		}
		d <<= 1
	}
	return r
}
