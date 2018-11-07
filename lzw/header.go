package lzw

const HeadLen = 30

func getHeader(src []byte) []byte {
    return src[:HeadLen]
}
