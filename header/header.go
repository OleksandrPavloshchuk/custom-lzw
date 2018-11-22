package header

import (
	"../version"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
)

var signature = []byte{0xAA, 'r', 0xCC}

const (
	unpackedSizeOffset = 6
	packedSizeOffset   = 14
	unpackedCrcOffset  = 22
	packedCrcOffset    = 38
)

var header [54]byte

func GetLength() int {
	return len(header)
}

func Fill(src *[]byte) {
	for i := 0; i < len(header); i++ {
		header[i] = (*src)[i]
	}
}

func AddHeader(src *[]byte) *[]byte {
	r := make([]byte, 0)
	r = append(r, header[:]...)
	r = append(r, *src...)
	return &r
}

func Print(h []byte) {
	fmt.Printf("Archive header:\n")
	fmt.Printf("- signature:     ")
	printHex(h[:3])
	fmt.Printf("\n- version:       %v.%v.%v\n", h[3], h[4], h[5])

	nu := uint64(binary.LittleEndian.Uint64(h[unpackedSizeOffset:packedSizeOffset]))
	unpackedSize := toString(nu)
	np := uint64(binary.LittleEndian.Uint64(h[packedSizeOffset:unpackedCrcOffset]))
	packedSize := toString(np)

	fieldWidth := len(unpackedSize)
	if len(packedSize) > fieldWidth {
		fieldWidth = len(packedSize)
	}
	f := fmt.Sprintf("- unpacked size: %v%vs\n", "%", fieldWidth)
	fmt.Printf(f, unpackedSize)
	f = fmt.Sprintf("- packed size:   %s%vs (%s.2f%s)\n", "%", fieldWidth, "%", "%s")
	fmt.Printf(f, packedSize, float32(np)/float32(nu)*100.0, "%")

	fmt.Printf("- unpacked CRC:  ")
	printHex(h[unpackedCrcOffset:packedCrcOffset])
	fmt.Printf("\n- packed CRC:    ")
	printHex(h[packedCrcOffset:len(header)])
	fmt.Printf("\n")
}

func CheckPackedContent(src *[]byte) error {
	if !checkSignature() {
		return errors.New("invalid archive signature")
	}
	if !checkVersion() {
		return errors.New("invalid archive version")
	}
	if !checkPackedSize(src) {
		return errors.New("invalid packed content size")
	}
	if !checkPackedCRC(src) {
		return errors.New("invalid packed CRC")
	}
	return nil
}

func CheckUnpackedContent(src *[]byte) error {
	if !checkUnpackedSize(src) {
		return errors.New("invalid unpacked content size")
	}
	if !checkUnpackedCRC(src) {
		return errors.New("invalid unpacked content CRC")
	}
	return nil
}

func SetSignature() {
	setArea(0, signature)
}

func checkSignature() bool {
	return checkArea(0, signature)
}

func SetVersion() {
	setArea(len(signature), version.ForHeader())
}

func checkVersion() bool {
	s := header[:]
	return version.IsCorrect(len(signature), &s)
}

func SetUnpackedInfo(src *[]byte) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(len(*src)))
	setArea(unpackedSizeOffset, b)
	s := md5.Sum(*src)
	setArea(unpackedCrcOffset, s[:])
}

func SetPackedInfo(src *[]byte) {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(len(*src)))
	setArea(packedSizeOffset, b)
	s := md5.Sum(*src)
	setArea(packedCrcOffset, s[:])
}

func checkUnpackedSize(src *[]byte) bool {
	toCheck := uint64(binary.LittleEndian.Uint64(header[unpackedSizeOffset : unpackedSizeOffset+8]))
	return uint64(len(*src)) == toCheck
}

func checkPackedSize(src *[]byte) bool {
	size := uint64(binary.LittleEndian.Uint64(header[packedSizeOffset : packedSizeOffset+8]))
	return uint64(len(*src)) == size
}

func checkUnpackedCRC(src *[]byte) bool {
	s := md5.Sum(*src)
	return checkArea(unpackedCrcOffset, s[:])
}

func checkPackedCRC(src *[]byte) bool {
	s := md5.Sum(*src)
	return checkArea(packedCrcOffset, s[:])
}

func setArea(offset int, src []byte) {
	for i := 0; i < len(src); i++ {
		header[offset+i] = src[i]
	}
}

func checkArea(offset int, src []byte) bool {
	for i := 0; i < len(src); i++ {
		if header[offset+i] != src[i] {
			return false
		}
	}
	return true
}

func toString(n uint64) string {
	src := []byte(fmt.Sprintf("%v", n))
	r := ""
	for i := len(src) - 1; i >= 0; i-- {
		r = string(src[i]) + r
		if (len(src)-i)%3 == 0 {
			r = " " + r
		}
	}
	return strings.Trim(r, " ")
}

func printHex(b []byte) {
	for _, v := range b {
		fmt.Printf("%02x ", v)
	}
}
