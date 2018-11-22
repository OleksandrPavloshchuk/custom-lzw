package codesIO

const size = 64

type bits uint64

type bitSet struct {
	words        []bits
	maxUsedIndex uint
}

func (this *bitSet) set(i uint) {
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

func (this *bitSet) isSet(i uint) bool {
    s := i/size
    if s >= uint(len(this.words)) {
        return false
    }
	return this.words[s]&(1<<(i%size)) != 0
}

func (this *bitSet) length() uint {
	return this.maxUsedIndex
}
