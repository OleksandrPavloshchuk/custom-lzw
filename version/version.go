package version

import (
	"fmt"
	"io"
)

const major = byte(0)
const minor = byte(0)
const patch = byte(3)
const date = "2018-11-07"

func Print(writer io.Writer) {
	fmt.Fprintf(writer, "Version: %d.%d.%d %v\n", major, minor, patch, date)
}

func IsCorrect(v []byte) bool {
    return isMajorCorrect(v[0]) && isMinorCorrect(v[1]) && isPatchCorrect(v[2])
}

func isMajorCorrect(v byte) bool {
    return v==0
}

func isMinorCorrect(v byte) bool {
    return v==0
}

func isPatchCorrect(v byte) bool {
    return v==3
}
