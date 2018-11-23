package static

import (
	//    "../../codesIO"
	"../../header"
	"fmt"
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

	contentOffset := header.GetLength() + int(header.GetCodeTableLength())
	
	fmt.Printf("TRACE content offset = %v\n", contentOffset)
	
	content := 	(*src)[contentOffset:]
	
	if err := header.CheckPackedContent(&content); err != nil {
		return nil, err
	}

	// TODO
	res := &content

	return res, nil
}
