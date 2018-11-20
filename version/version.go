package version

import (
	"fmt"
)

const major = byte(0)
const minor = byte(0)
const patch = byte(6)
const date = "2018-11-20"

func ForHeader() []byte {
    return []byte{major,minor,patch}
}

func Print() {
	fmt.Printf("Version: %d.%d.%d %v\n", major, minor, patch, date)
}

func IsCorrect(offset int, v *[]byte) bool {
    return isMajorCorrect(get(offset, v, 0)) && isMinorCorrect(get(offset, v, 1)) && isPatchCorrect(get(offset, v, 2))
}

func get(offset int, v *[]byte, i int) byte {
    return (*v)[offset+i]
}

func isMajorCorrect(v byte) bool {
    return v==major
}

func isMinorCorrect(v byte) bool {
    return v==minor
}

func isPatchCorrect(v byte) bool {
    return v==patch
}
