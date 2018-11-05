package main

import (
	"io/ioutil"
)

type CodeWriter struct {
	start uint
        bitSet BitSet	
}

func (this *CodeWriter) Accept(code uint, length uint) {
	d := uint(1)
	for i:=uint(0); i<length; i++ {
		if code & d != 0 {
			this.bitSet.Set(this.start)
		}
		this.start++
		d <<= 1		
	}
}

func (this *CodeWriter) Write(fileName string) error {
	result := make([]byte, (this.start + 7) / 8)
	for i:=0; i<len(result); i++ {
		result[i] = this.toByte(uint(i << 3))
	}
	return ioutil.WriteFile(fileName, result, 0644)
}

func (this *CodeWriter) toByte(offset uint) byte {
	r := byte(0)
	d := byte(1)
	for i := 0; i<8; i++ {
		if this.bitSet.IsSet(uint(i)+offset) {
			r |= d
		}
		d <<= 1
	}	
	return r
}



