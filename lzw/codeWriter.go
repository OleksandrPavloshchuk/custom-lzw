package lzw

type codeWriter struct {
	codeIO
}

func (this *codeWriter) accept(code uint, length uint) {
	d := uint(1)
	for i := uint(0); i < length; i++ {
		if code&d != 0 {
			this.bitSet.set(this.start)
		}
		this.start++
		d <<= 1
	}
}

func (this *codeWriter) getBytes() []byte {
	result := make([]byte, HeadLen)
	if 0 == this.start {
		return result
	}
	size := (this.start + 7) / 8
	for i := 0; uint(i) < size; i++ {
		result = append(result, this.toByte(uint(i<<3)))
	}
	return result
}

func (this *codeWriter) toByte(offset uint) byte {
	r := byte(0)
	d := byte(1)
	for i := 0; i < 8; i++ {
		if this.bitSet.isSet(uint(i) + offset) {
			r |= d
		}
		d <<= 1
	}
	return r
}
