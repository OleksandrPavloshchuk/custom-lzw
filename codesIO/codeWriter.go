package codesIO

type Writer struct {
	codeIO
}

func (cw *Writer) Accept(code uint, length uint) {
	d := uint(1)
	for i := uint(0); i < length; i++ {
		if code&d != 0 {
			cw.bitSet.set(cw.start)
		}
		cw.start++
		d <<= 1
	}
}

func (cw *Writer) GetBytes() []byte {
	result := make([]byte,0)
	if 0 == cw.start {
		return result
	}
	size := (cw.start + 7) / 8
	for i := 0; uint(i) < size; i++ {
		result = append(result, cw.toByte(uint(i<<3)))
	}
	return result
}

func (cw *Writer) toByte(offset uint) byte {
	r := byte(0)
	d := byte(1)
	for i := 0; i < 8; i++ {
		if cw.bitSet.isSet(uint(i) + offset) {
			r |= d
		}
		d <<= 1
	}
	return r
}
