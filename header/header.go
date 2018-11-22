package header

import (
	"../version"
	"crypto/md5"
	"errors"
	"encoding/binary"
)

const Length = 54

var signature = []byte{0xAA, 'r', 0xCC}

const (
	unpackedSizeOffset = 6
	packedSizeOffset   = 14
	unpackedCrcOffset  = 22
	packedCrcOffset    = 38
)

type Header struct {
	buf *[]byte
}

func GetHeader(src *[]byte) Header {
    return Header{buf:src}
}

func (h *Header) CheckPackedContent() error {
	if !h.checkSignature() {
		return errors.New("invalid archive signature")
	}
	if !h.checkVersion() {
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

func (h *Header) CheckUnpackedContent(res *[]byte) error {
	if !h.checkUnpackedSize(uint64(len(*res))) {
		return errors.New("invalid unpacked content size")
	}
	if !h.checkUnpackedCRC(res) {
		return errors.New("invalid unpacked content CRC")
	}
	return nil
}

func SetHeader(res *[]byte, src *[]byte) {
	h := Header{res}
	h.setSignature()
	h.setVersion()
	h.setUnpackedSize(uint64(len(*src)))
	h.setPackedSize()
	h.setUnpackedCRC(src)
	h.setPackedCRC()
}

func (h *Header) setSignature() {
	h.setArea(0, signature)
}

func (h *Header) checkSignature() bool {
	return h.checkArea(0, signature)
}

func (h *Header) setVersion() {
	h.setArea(len(signature), version.ForHeader())
}

func (h *Header) checkVersion() bool {
	return version.IsCorrect(len(signature), h.buf)
}

func (h *Header) setUnpackedSize(size uint64) {
    b := make([]byte,8)
    binary.LittleEndian.PutUint64(b, size)
	h.setArea(unpackedSizeOffset, b)
}

func (h *Header) setPackedSize() {
    b := make([]byte,8)
    binary.LittleEndian.PutUint64(b, uint64(len(*h.buf)-Length))
	h.setArea(packedSizeOffset, b)
}

func (h *Header) checkUnpackedSize(size uint64) bool {
    toCheck := uint64(binary.LittleEndian.Uint64((*h.buf)[unpackedSizeOffset:unpackedSizeOffset+8]))
	return size == toCheck
}

func (h *Header) checkPackedSize() bool {
	size := uint64(binary.LittleEndian.Uint64((*h.buf)[packedSizeOffset : packedSizeOffset+8]))
	return uint64(len(*h.buf)-Length) == size
}

func (h *Header) setUnpackedCRC(src *[]byte) {
	s := md5.Sum(*src)
	h.setArea(unpackedCrcOffset, s[:])
}

func (h *Header) checkUnpackedCRC(src *[]byte) bool {
	s := md5.Sum(*src)
	return h.checkArea(unpackedCrcOffset, s[:])
}

func (h *Header) setPackedCRC() {
	s := md5.Sum((*h.buf)[Length:])
	h.setArea(packedCrcOffset, s[:])
}

func (h *Header) checkPackedCRC() bool {
	s := md5.Sum((*h.buf)[Length:])
	return h.checkArea(packedCrcOffset, s[:])
}

func (h *Header) setArea(offset int, src []byte) {
	for i := 0; i < len(src); i++ {
		(*h.buf)[offset+i] = src[i]
	}
}

func (h *Header) checkArea(offset int, src []byte) bool {
	for i := 0; i < len(src); i++ {
		if (*h.buf)[offset+i] != src[i] {
			return false
		}
	}
	return true
}
