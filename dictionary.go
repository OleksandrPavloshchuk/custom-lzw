package main

const IncrementCodeSize = 0

type Dictionary struct {
    data [][]byte
    codeSize uint
}

func (this *Dictionary) Init() {
    this.codeSize = 9
    for b:=0; b<256; b++ {
        var a [1]byte
        a[0] = byte(b)
        this.Put(a[:])
    }    
}

func (this *Dictionary) GetCodeSize() uint {
    return this.codeSize
}

func (this *Dictionary) Put(s []byte) {
    this.data = append( this.data, s)
}

func (this *Dictionary) GetString(i uint) []byte {
    return this.data[i-1]
}

func (this *Dictionary) GetIndex(a []byte) uint {
    for i,s := range this.data {
        if areEqual(s, a) {
            return uint(i+1)
        }
    }
    return 0
}

func (this *Dictionary) IncrementCodeSizeWhileDecode(code uint) bool {
    return this.incrementCodeSizeWhenCondition( code == IncrementCodeSize )
}

func (this *Dictionary) IncrementCodeSizeWhileEncode() bool {
    return this.incrementCodeSizeWhenCondition( uint(len(this.data))+1 > (1<<this.codeSize))
}

func (this *Dictionary) HasString(a []byte) bool {
    for _,s := range this.data {
        if areEqual(s,a) {
            return true
        }
    }
    return false
}

func (this *Dictionary) HasCode(i uint) bool {
    return i==0 || i>uint(len(this.data))
}

func (this *Dictionary) incrementCodeSizeWhenCondition(condition bool) bool {
    if condition {
        this.codeSize++
        return true;
    }
    return false
}

func areEqual(a1 []byte, a2 []byte) bool {
    if len(a1)!=len(a2) {
        return false
    }
    for i,b:=range a1 {
        if b != a2[i] {
            return false
        }
    }
    return true
}

