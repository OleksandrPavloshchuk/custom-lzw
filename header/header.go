package header

import (
	"../version"
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"
	"strconv"
)

var signature = []byte{0xAA, 'r', 0xCC}

const (
	signatureOffset            = 0
	versionOffset              = 3
	unpackedSizeOffset         = 6
	packedSizeOffset           = 10
	unpackedCrcOffset          = 14
	packedCrcOffset            = 30
	codeTableLengthOffset      = 46
	
	formatUnpackedSizeStart    = "- unpacked size:     %"
	formatUnpackedSizeEnd      = "s\n"
	
	formatPackedSizeStart      = "- packed size:       %"
	formatPackedSizeEnd        = "s (%.2f%s)\n"
	
	formatCodeTableLengthStart = "- code table length: %"
	formatCodeTableLengthEnd   = "s\n"	 
	
)

var header [50]byte

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

func Print(h *[]byte) {
	Fill(h)

	fmt.Printf("Archive header:\n")
	fmt.Printf("- signature:         ")
	printHex(signatureOffset, versionOffset)
	fmt.Printf("- version:           %v.%v.%v\n", header[3], header[4], header[5])

	nu := toUint32(unpackedSizeOffset)
	unpackedSize := toString(nu)
	np := toUint32(packedSizeOffset)
	packedSize := toString(np)

	fieldWidth := len(unpackedSize)
	if len(packedSize) > fieldWidth {
		fieldWidth = len(packedSize)
	}
	fieldWidthStr := strconv.Itoa(fieldWidth)
	fmt.Printf(formatUnpackedSizeStart + fieldWidthStr + formatUnpackedSizeEnd, unpackedSize)
	fmt.Printf(formatPackedSizeStart + fieldWidthStr + formatPackedSizeEnd, packedSize, float32(np)/float32(nu)*100.0, "%" )

	fmt.Printf("- unpacked CRC:      ")
	printHex(unpackedCrcOffset, packedCrcOffset)
	fmt.Printf("- packed CRC:        ")
	printHex(packedCrcOffset, codeTableLengthOffset)
	fmt.Printf(formatCodeTableLengthStart + fieldWidthStr + formatCodeTableLengthEnd, toString(GetCodeTableLength()))
}

func GetCodeTableLength() uint32 {
    return toUint32(codeTableLengthOffset)
}

func SetCodeTableLength(length uint32) {
    toBytes(length, codeTableLengthOffset)
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

func setInfo(src *[]byte, sizeOffset int, crcOffset int) {
	toBytes(getLengthUint32(src), sizeOffset)
	s := md5.Sum(*src)
	setArea(crcOffset, s[:])
}

func SetUnpackedInfo(src *[]byte) {
	setInfo(src, unpackedSizeOffset, unpackedCrcOffset)
}

func SetPackedInfo(src *[]byte) {
	setInfo(src, packedSizeOffset, packedCrcOffset)
}

func checkSize(src *[]byte, offset int) bool {
	return getLengthUint32(src) == toUint32(offset)
}

func checkUnpackedSize(src *[]byte) bool {
	return checkSize(src, unpackedSizeOffset)
}

func checkPackedSize(src *[]byte) bool {
	return checkSize(src, packedSizeOffset)
}

func checkCRC(src *[]byte, offset int) bool {
	s := md5.Sum(*src)
	return checkArea(offset, s[:])
}

func checkUnpackedCRC(src *[]byte) bool {
	return checkCRC(src, unpackedCrcOffset)
}

func checkPackedCRC(src *[]byte) bool {
	return checkCRC(src, packedCrcOffset)
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

func toUint32(start int) uint32 {
	return uint32(binary.LittleEndian.Uint32(header[start : start+4]))
}

func toBytes(n uint32, start int) {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, n)
	setArea(start, b)
}

func getLengthUint32(src *[]byte) uint32 {
	return uint32(len(*src))
}

func toString(n uint32) string {
	src := []byte(strconv.FormatUint(uint64(n), 10))	
	r := ""
	for i := len(src) - 1; i >= 0; i-- {
		r = string(src[i]) + r
		if (len(src)-i)%3 == 0 {
			r = " " + r
		}
	}
	return strings.Trim(r, " ")
}

func printHex(start int, end int) {
	for i := start; i < end; i++ {
		fmt.Printf("%02x ", header[i])
	}
	fmt.Printf("\n")
}
