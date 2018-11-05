package main

const size = 64

type bits uint64

// BitSet is a set of bits that can be set, cleared and queried.
type BitSet []bits

// Set ensures that the given bit is set in the BitSet.
func (this *BitSet) Set(i uint) {
    if len(*this) < int(i/size+1) {
        r := make([]bits, i/size+1)
        copy(r, *this)
        *this = r
    }
    (*this)[i/size] |= 1 << (i % size)
}

// IsSet returns true if the given bit is set, false if it is cleared.
func (this *BitSet) IsSet(i uint) bool {
    return (*this)[i/size]&(1<<(i%size)) != 0
}
