package main

import (
	"fmt"
	"io"
)

const major = 0
const minor = 0
const patch = 3
const date = "2018-11-07"

func PrintVersion(writer io.Writer) {
	fmt.Fprintf(writer, "Version: %d.%d.%d %v\n", major, minor, patch, date)
}
