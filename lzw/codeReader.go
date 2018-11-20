package lzw

type codeReader struct {
	codeIO
}

func (this *codeReader) hasCodes() bool {
	return this.start <= this.bitSet.Length()
}

func (this *codeReader) get(codeLength uint) uint {
	r := uint(0)
	d := uint(1)
	for i := uint(0); i < codeLength; i++ {
		if this.bitSet.isSet(this.start) {
			r |= d
		}
		this.start++
		d <<= 1
	}
	return r
}

func acquireCodes(src []byte) codeReader {
	cr := codeReader{}
	counter := uint(0)
	for n, b := range src {
		if n >= HeadLen {
			d := byte(1)
			for i := 0; i < 8; i++ {
				if d&b != 0 {
					cr.bitSet.set(counter)
				}
				counter++
				d <<= 1
			}
		}
	}
	return cr
}
