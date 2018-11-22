package static

/*
import (
    "../../codesIO"
    "../../header"
)
*/

// TODO replace it by static Huffman's encoding

/*
func decode(cr codesIO.Reader) []byte {
    return make([]byte,0)
}
*/

func Decode(src []byte, _ uint64) ([]byte, error) {    
	if len(src) == 0 {
		return []byte{}, nil
	}
	// TODO
	return src, nil
}
