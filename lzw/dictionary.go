package lzw

const IncrementCodeSize = 0

type dictionary struct {
	data     [][]byte
	index    map[string]uint
	codeSize uint
}

func createDictionary() dictionary {
	d := dictionary{index: make(map[string]uint), codeSize: 9}
	for b := 0; b < 256; b++ {
		var a [1]byte
		a[0] = byte(b)
		d.put(a[:])
	}
	return d
}

func (this *dictionary) getCodeSize() uint {
	return this.codeSize
}

func (this *dictionary) put(s []byte) {
	this.data = append(this.data, s)
	this.index[string(s)] = uint(len(this.data))
}

func (this *dictionary) getString(i uint) []byte {
	return this.data[i-1]
}

func (this *dictionary) getIndex(a []byte) uint {
	return this.index[string(a)]
}

func (this *dictionary) incrementCodeSizeWhileDecode(code uint) bool {
	return this.incrementCodeSizeWhenCondition(code == IncrementCodeSize)
}

func (this *dictionary) incrementCodeSizeWhileEncode() bool {
	return this.incrementCodeSizeWhenCondition(uint(len(this.data))+1 > (1 << this.codeSize))
}

func (this *dictionary) hasString(a []byte) bool {
	_, r := this.index[string(a)]
	return r
}

func (this *dictionary) hasCode(i uint) bool {
	return i != 0 && i <= uint(len(this.data))
}

func (this *dictionary) incrementCodeSizeWhenCondition(condition bool) bool {
	if condition {
		this.codeSize++
		return true
	}
	return false
}
