package static

import (
	//    "../../codesIO"
	"../../header"
)

// TODO replace it by static Huffman's encoding

/*
func decode(cr codesIO.Reader) []byte {
    return make([]byte,0)
}
*/

func Decode(src *[]byte) (*[]byte, error) {
	if len(*src) == 0 {
		res := []byte{}
		return &res, nil
	}
	header.Fill(src)

	content := (*src)[header.GetLength():]

	if err := header.CheckPackedContent(&content); err != nil {
		return nil, err
	}

	// TODO
	res := &content

	return res, nil
}
