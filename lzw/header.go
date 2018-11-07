package lzw

import (
    "crypto/md5"
)

const HeadLen = 54

var signature = []byte{0xAA, 'r', 0xCC}
const unpackedSizeOffset = 6
const packedSizeOffset = 14
const unpackedCrcOffset = 22
const packedCrcOffset = 38

type Header struct {
    buf *[]byte
}

func (this *Header) SetSignature() {
    this.setArea(0, signature)
}

func (this *Header) CheckSignature() bool {
    return this.checkArea(0, signature)
}

func (this *Header) SetVersion(src []byte) {
    this.setArea(len(signature), src)
}

func (this *Header) CheckVersion(checker func(int,*[]byte) bool) bool {
    return checker(len(signature), this.buf)
}

func (this *Header) SetUnpackedSize(size uint64) {
    this.setArea( unpackedSizeOffset, toBytes(size))
}

func (this *Header) SetPackedSize() {
    this.setArea( packedSizeOffset, toBytes(uint64(len(*this.buf)-HeadLen)))
}

func (this *Header) CheckUnpackedSize(size uint64) bool {
    return size==fromBytes((*this.buf)[unpackedSizeOffset:unpackedSizeOffset+8])
}

func (this *Header) CheckPackedSize() bool {
    return uint64(len(*this.buf)-HeadLen)==fromBytes((*this.buf)[packedSizeOffset:packedSizeOffset+8])
}

func (this *Header) SetUnpackedCRC(src *[]byte) {
    s := md5.Sum(*src)
    this.setArea( unpackedCrcOffset, s[:])
}

func (this *Header) CheckUnpackedCRC(src *[]byte) bool {
    s := md5.Sum(*src)
    return this.checkArea(unpackedCrcOffset, s[:]);
}

func (this *Header) SetPackedCRC() {
    s := md5.Sum((*this.buf)[HeadLen:])
    this.setArea( packedCrcOffset, s[:])
}

func (this *Header) CheckPackedCRC() bool {
    s := md5.Sum((*this.buf)[HeadLen:])
    return this.checkArea(packedCrcOffset, s[:]);
}

func (this *Header) setArea(offset int, src []byte) {
	for i := 0; i < len(src); i++ {
		(*this.buf)[offset+i] = src[i]
	}    
}

func (this *Header) checkArea(offset int, src []byte) bool {
	for i := 0; i < len(src); i++ {
		if (*this.buf)[offset+i] != src[i] {
			return false
		}
	}
	return true
}

func toBytes(n uint64) []byte {
    var r [8]byte
    for i:=7; i>=0; i-- {
        r[i] = byte(0xFF & (n >> uint(8*(7-i)) ) )
    }
    return r[:]
}

func fromBytes(src []byte) uint64 {
    r:=uint64(0)
    for i:=0; i<8; i++ {
        r <<= 8
        r |= uint64(src[i])        
    }    
    return r
}

