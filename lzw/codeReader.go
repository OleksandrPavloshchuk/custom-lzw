package lzw

import (
	"io/ioutil"
)

type CodeReader struct {
	CodeIO
}

func (this *CodeReader) HasCodes() bool {
	return this.start <= this.bitSet.Length()
}

func (this *CodeReader) Get(codeLength uint) uint {
	r := uint(0)
	d := uint(1)
	for i := uint(0); i < codeLength; i++ {
		if this.bitSet.IsSet(this.start) {
			r |= d
		}
		this.start++
		d <<= 1
	}
	return r
}

func (this *CodeReader) Read(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	counter := uint(0)
	for n, b := range data {
	    if n >= HeadLen {
    		d := byte(1)
	    	for i := 0; i < 8; i++ {
	    		if d&b != 0 {
	    			this.bitSet.Set(counter)
	    		}
	    		counter++
	    		d <<= 1
	    	}
	    }
	}
	return nil
}
