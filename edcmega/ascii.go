package edcmega

import (
	"fmt"
	"strings"
)

// https://www.geeksforgeeks.org/what-is-ascii-a-complete-guide-to-generating-ascii-code/?ref=lbp

func ToHex(str string) string {
	res := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		d := str[i]
		h := fmt.Sprintf("%x", d)
		res[i] = h
	}
	return strings.Join(res[:], "")
}
