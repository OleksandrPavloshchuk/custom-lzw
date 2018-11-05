package main

import (
//	"fmt"
)

func writeTestData(data []uint, codeSize uint, fileName string) error {
	codeWriter := CodeWriter{}
	for _, i := range data {
		codeWriter.Accept(i, codeSize)
	}
	return codeWriter.Write(fileName)
}


func main() {
	const codeSize = 9
	srcData := []uint {1, 1, 1, 1}
	const fileName = "/home/pavloshchuk-ov/temp/test1/1.A"

	err := writeTestData(srcData, codeSize, fileName)
	if err!=nil {
		panic(err)
	}
}
