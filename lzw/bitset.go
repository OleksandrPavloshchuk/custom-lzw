package lzw

const size = 64

type bits uint64

type BitSet struct {
	words        []bits
	maxUsedIndex uint
}

func (this *BitSet) Set(i uint) {
	if this.maxUsedIndex < i {
		this.maxUsedIndex = i
	}
	newSize := int(i/size + 1)

	if len(this.words) < newSize {
		n := make([]bits, newSize+2)
		copy(n, this.words)
		this.words = n
	}
	this.words[i/size] |= 1 << (i % size)
}

func (this *BitSet) IsSet(i uint) bool {
	return this.words[i/size]&(1<<(i%size)) != 0
}

func (this *BitSet) Length() uint {
	return this.maxUsedIndex
}
