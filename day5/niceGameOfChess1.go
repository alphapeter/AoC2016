package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	input := "cxdnnyjw"

	l := 0
	for i := 0; l < 8; i++ {
		h := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		hash := fmt.Sprintf("%x", h)

		if hash[0:5] == "00000" {
			fmt.Printf("%c", hash[5])
			l++
		}
	}
}
