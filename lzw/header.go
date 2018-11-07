package lzw

const HeadLen = 30
var signature = []byte{ 0xAA, 'r', 0xCC }

func SetSignature(head *[]byte) {
    for i:=0; i<len(signature); i++ {
        (*head)[i] = signature[i]
    }
}

func CheckSignature(head *[]byte) bool {
    for i:=0; i<len(signature); i++ {
        if (*head)[i] != signature[i] {
            return false
        }
    }
    return true
}
