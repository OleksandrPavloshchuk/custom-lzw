package lzw

import  (
	"fmt"
)

const IncrementCodeSize = 0

type dictionary struct {
	data     [][]byte
    weight []uint64
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
    this.weight = append(this.weight, 0)
}

func (this *dictionary) getString(i uint) []byte {
    this.weight[i-1]++
	return this.data[i-1]
}

func (this *dictionary) getIndex(a []byte) uint {
    r := this.index[string(a)]
    this.weight[r-1]++
    return r
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

func (this *dictionary) PrintStatistics() {
	fmt.Printf("Archive statistics:\n")	
	fmt.Printf("\tcode size: %v\n", this.codeSize)	
	fmt.Printf("\tdictionary size: %v\n", len(this.data))
	fmt.Printf("\tdictionary\n",)
	fmt.Printf("\t\tindex\tweight\tvalue\n",)
	fmt.Printf("\t\t-----\t------\t-----\n",)
   for i, v := range this.data {
        w := this.weight[i]
		fmt.Printf("\t\t%5d\t%6d\t%v\n",  i, w, v)
   }
	fmt.Printf("\n",)
}
