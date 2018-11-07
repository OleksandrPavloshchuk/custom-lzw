package main

const IncrementCodeSize = 0

type Dictionary struct {
	data     [][]byte
	index    map[string]uint
	codeSize uint
}

func (this *Dictionary) Init() {
	this.index = make(map[string]uint)
	this.codeSize = 9
	for b := 0; b < 256; b++ {
		var a [1]byte
		a[0] = byte(b)
		this.Put(a[:])
	}
}

func (this *Dictionary) GetCodeSize() uint {
	return this.codeSize
}

func (this *Dictionary) Put(s []byte) {
	this.data = append(this.data, s)
	this.index[string(s)] = uint(len(this.data))
}

func (this *Dictionary) GetString(i uint) []byte {
	return this.data[i-1]
}

func (this *Dictionary) GetIndex(a []byte) uint {
	return this.index[string(a)]
}

func (this *Dictionary) IncrementCodeSizeWhileDecode(code uint) bool {
	return this.incrementCodeSizeWhenCondition(code == IncrementCodeSize)
}

func (this *Dictionary) IncrementCodeSizeWhileEncode() bool {
	return this.incrementCodeSizeWhenCondition(uint(len(this.data))+1 > (1 << this.codeSize))
}

func (this *Dictionary) HasString(a []byte) bool {
	_, r := this.index[string(a)]
	return r
}

func (this *Dictionary) HasCode(i uint) bool {
	return i != 0 && i <= uint(len(this.data))
}

func (this *Dictionary) incrementCodeSizeWhenCondition(condition bool) bool {
	if condition {
		this.codeSize++
		return true
	}
	return false
}
