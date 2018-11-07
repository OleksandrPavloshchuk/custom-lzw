package version

import (
	"fmt"
	"io"
)

const major = byte(0)
const minor = byte(0)
const patch = byte(3)
const date = "2018-11-07"

func ForHeader() []byte {
    return []byte{major,minor,patch}
}

func Print(writer io.Writer) {
	fmt.Fprintf(writer, "Version: %d.%d.%d %v\n", major, minor, patch, date)
}

func IsCorrect(offset int, v *[]byte) bool {
    return isMajorCorrect(get(offset, v, 0)) && isMinorCorrect(get(offset, v, 1)) && isPatchCorrect(get(offset, v, 2))
}

func get(offset int, v *[]byte, i int) byte {
    return (*v)[offset+i]
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
