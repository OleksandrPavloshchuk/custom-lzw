package codesIO

type CodeReader struct {
	codeIO
}

func (cr *CodeReader) HasCodes() bool {
	return cr.start <= cr.bitSet.length()
}

func (cr *CodeReader) Get(codeLength uint) uint {
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

func AcquireCodes(src []byte, offset int) CodeReader {
	cr := CodeReader{}
	counter := uint(0)
	for n, b := range src {
		if n >= offset {
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
