package static

import (
	//    "../../codesIO"
	"../../header"
	"encoding/binary"	
	
	"fmt"
)

// TODO replace it by static Huffman's encoding

/*
func encode(src []byte, cw *codesIO.Writer) {
	if len(src) == 0 {
		return
	}
	// TODO
}

func emit(s []byte, dict dictionary, cw *codesIO.Writer) {
	cw.Accept(dict.getIndex(s), dict.getCodeSize())
}
*/

func serializeItem(key string, value uint32, offset int, dest *[]byte) {
	(*dest)[offset] = []byte(key)[0]
	binary.LittleEndian.PutUint32((*dest)[offset+1:offset+5], value)
}

func serialize(src *map[string]uint32) []byte {
    r := make([]byte, len(*src)*5)
    i := 0
    for key, value := range *src {
        serializeItem( key, value, i, &r)
        i += 5
    }    
    return r
}

func Encode(src *[]byte) (*[]byte, error) {

    hTree := BuildHuffmanTree(src)
    codes := hTree.GetCodes()
    
    codeTableBytes := serialize(&codes)
    
    fmt.Printf("TRACE: code table=%v\nTRACE serialized=%v\n", codes, codeTableBytes)
    
    
    res := make([]byte, 0)
    res = append( res, codeTableBytes...)
    
	// TODO:
	
	
	res = append(res, *src...)

    header.SetCodeTableLength(uint32(len(codeTableBytes)))
	header.SetPackedInfo(&res)
	return &res, nil
}
