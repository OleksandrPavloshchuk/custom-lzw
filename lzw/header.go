package lzw

const HeadLen = 30

var signature = []byte{0xAA, 'r', 0xCC}

func SetSignature(head *[]byte) {
    setArea(head, 0, signature)
}

func CheckSignature(head *[]byte) bool {
    return checkArea(head, 0, signature)
}

func SetVersion(head *[]byte, src []byte) {
    setArea(head, len(signature), src)
}

func CheckVersion(head *[]byte, src []byte) bool {
    return checkArea(head, len(signature), src)
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

