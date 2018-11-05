package main

import (
	"fmt"
)

func writeTestData(data []uint, codeSize uint, fileName string) error {
	codeWriter := CodeWriter{}
	for _, i := range data {
		codeWriter.Accept(i, codeSize)
	}
	return codeWriter.Write(fileName)
}

func readTestData(codeSize uint, fileName string) ([]uint, error) {
	codeReader := CodeReader{}
	err := codeReader.Read(fileName)
	if err!=nil {
	    return nil, err
	}
	r := make([]uint, 0)
	for codeReader.HasCodes() {
	    r = append(r, codeReader.Get(codeSize))
	}
	
	return r, nil
}



func main() {
	const codeSize = 9
	srcData := []uint {1, 1, 1, 1}
	const fileName = "/home/pavloshchuk-ov/temp/test1/1.A"

	err := writeTestData(srcData, codeSize, fileName)
	if err!=nil {
		panic(err)
	}
	
	resData,err := readTestData(codeSize, fileName)
	if err!=nil {
	    panic(err)
	}
	
    fmt.Printf("SRC=%v\nRES=%v\n", srcData, resData)	
}
