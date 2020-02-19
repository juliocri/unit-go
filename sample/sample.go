package sample

import (
	"errors"
	"fmt"
	"strconv"
)

// CidrToMask - Return the right net mask given by a CIDR integer.
// Example: 32 => 255.255.255.255
func CidrToMask(c int) (string, error) {
	if c < 1 || c > 32 {
		return "", errors.New("cidr: not valid number")
	}
	var b [32]int
	var s = ""
	for i := 0; i < 32; i++ {
		if i < c {
			b[i] = 1
		}
		s += strconv.Itoa(b[i])
	}

	o1, _ := strconv.ParseInt(s[:8], 2, 64)
	o2, _ := strconv.ParseInt(s[8:16], 2, 64)
	o3, _ := strconv.ParseInt(s[16:24], 2, 64)
	o4, _ := strconv.ParseInt(s[24:], 2, 64)

	return fmt.Sprintf("%v.%v.%v.%v", o1, o2, o3, o4), nil
}
