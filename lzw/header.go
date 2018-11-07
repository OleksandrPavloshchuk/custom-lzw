package lzw

const HeadLen = 30

var signature = []byte{0xAA, 'r', 0xCC}
const unpackedSizeOffset = 6
const packedSizeOffset = 14
const crcOffset = 22

func SetSignature(head *[]byte) {
    setArea(head, 0, signature)
}

func CheckSignature(head *[]byte) bool {
    return checkArea(head, 0, signature)
}

func SetVersion(head *[]byte, src []byte) {
    setArea(head, len(signature), src)
}

func CheckVersion(head *[]byte, checker func(int,*[]byte) bool) bool {
    return checker(len(signature), head)
}

func SetUnpackedSize(head *[]byte, size uint64) {
    setArea( head, unpackedSizeOffset, toBytes(size))
}

func SetPackedSize(head *[]byte) {
    setArea( head, packedSizeOffset, toBytes(uint64(len(*head)-HeadLen)))
}

func CheckUnpackedSize(head *[]byte, size uint64) bool {
    return size==fromBytes((*head)[unpackedSizeOffset:unpackedSizeOffset+8])
}

func CheckPackedSize(head *[]byte, size uint64) bool {
    return size-HeadLen==fromBytes((*head)[packedSizeOffset:packedSizeOffset+8])
}

func setArea(head *[]byte, offset int, src []byte) {
	for i := 0; i < len(src); i++ {
		(*head)[offset+i] = src[i]
	}    
}

func checkArea(head *[]byte, offset int, src []byte) bool {
	for i := 0; i < len(src); i++ {
		if (*head)[offset+i] != src[i] {
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

