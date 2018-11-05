package main

const size = 64

type bits uint64

type BitSet struct {
    words []bits
    maxUsedIndex uint
}

func (this *BitSet) Set(i uint) {
    if this.maxUsedIndex < i {
        this.maxUsedIndex = i
    }
    if len(this.words) < int(i/size+1) {
        r := BitSet{make([]bits, i/size+1), this.maxUsedIndex}
        copy(r.words, this.words)
        *this = r
    }
    this.words[i/size] |= 1 << (i % size)
}

func (this *BitSet) IsSet(i uint) bool {
    if i>this.maxUsedIndex {
        return false
    }
    return this.words[i/size]&(1<<(i%size)) != 0
}

func (this *BitSet) Length() uint {
	return this.maxUsedIndex
}
