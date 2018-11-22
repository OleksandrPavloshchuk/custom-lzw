package codesIO

type Reader struct {
	codeIO
}

func (cr *Reader) HasCodes() bool {
	return cr.start <= cr.bitSet.length()
}

func (cr *Reader) Get(codeLength uint) uint {
	r := uint(0)
	d := uint(1)
	for i := uint(0); i < codeLength; i++ {
		if cr.bitSet.isSet(cr.start) {
			r |= d
		}
		cr.start++
		d <<= 1
	}
	return r
}

func AcquireCodes(src *[]byte) Reader {
	cr := Reader{}
	counter := uint(0)
	for _, b := range *src {
		d := byte(1)
		for i := 0; i < 8; i++ {
			if d&b != 0 {
				cr.bitSet.set(counter)
			}
			counter++
			d <<= 1
		}
	}
	return cr
}
