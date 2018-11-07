package lzw

import (
    "crypto/md5"
    "errors"
)

const HeadLen = 54

var signature = []byte{0xAA, 'r', 0xCC}
const unpackedSizeOffset = 6
const packedSizeOffset = 14
const unpackedCrcOffset = 22
const packedCrcOffset = 38

type header struct {
    buf *[]byte
}

func (h *header) CheckPackedContent(version []byte) error {
	if !h.checkSignature() {
		return errors.New("invalid archive signature")
	}
	if !h.checkVersion(VersionChecker) {
		return errors.New("invalid archive version")
	}
	if !h.checkPackedSize() {
	    return errors.New("invalid packed content size")
	}
	if !h.checkPackedCRC() {
	    return errors.New("invalid packed CRC")
	}
	return nil
}

func (h *header) CheckUnpackedContent(res *[]byte) error {
	if !h.checkUnpackedSize(uint64(len(*res))) {
	    return errors.New("invalid unpacked content size")
	}
	if !h.checkUnpackedCRC(res) {
	    return errors.New("invalid unpacked content CRC")
	}
	return nil
}

func setHeader(res *[]byte, src *[]byte, version []byte) {
    h:=header{res}
	h.setSignature()
	h.setVersion(version)
	h.setUnpackedSize(uint64(len(*src)))
	h.setPackedSize()
	h.setUnpackedCRC(src)
	h.setPackedCRC()
}


func (h *header) setSignature() {
    h.setArea(0, signature)
}

func (h *header) checkSignature() bool {
    return h.checkArea(0, signature)
}

func (h *header) setVersion(src []byte) {
    h.setArea(len(signature), src)
}

func (h *header) checkVersion(checker func(int,*[]byte) bool) bool {
    return checker(len(signature), h.buf)
}

func (h *header) setUnpackedSize(size uint64) {
    h.setArea( unpackedSizeOffset, toBytes(size))
}

func (h *header) setPackedSize() {
    h.setArea( packedSizeOffset, toBytes(uint64(len(*h.buf)-HeadLen)))
}

func (h *header) checkUnpackedSize(size uint64) bool {
    return size==fromBytes((*h.buf)[unpackedSizeOffset:unpackedSizeOffset+8])
}

func (h *header) checkPackedSize() bool {
    return uint64(len(*h.buf)-HeadLen)==fromBytes((*h.buf)[packedSizeOffset:packedSizeOffset+8])
}

func (h *header) setUnpackedCRC(src *[]byte) {
    s := md5.Sum(*src)
    h.setArea( unpackedCrcOffset, s[:])
}

func (h *header) checkUnpackedCRC(src *[]byte) bool {
    s := md5.Sum(*src)
    return h.checkArea(unpackedCrcOffset, s[:]);
}

func (h *header) setPackedCRC() {
    s := md5.Sum((*h.buf)[HeadLen:])
    h.setArea( packedCrcOffset, s[:])
}

func (h *header) checkPackedCRC() bool {
    s := md5.Sum((*h.buf)[HeadLen:])
    return h.checkArea(packedCrcOffset, s[:]);
}

func (h *header) setArea(offset int, src []byte) {
	for i := 0; i < len(src); i++ {
		(*h.buf)[offset+i] = src[i]
	}    
}

func (h *header) checkArea(offset int, src []byte) bool {
	for i := 0; i < len(src); i++ {
		if (*h.buf)[offset+i] != src[i] {
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

