package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	input := "cxdnnyjw"

	password := []rune{'_', '_', '_', '_', '_', '_', '_', '_'}
	l := 0
	for i := 0; l < 8; i++ {
		h := md5.Sum([]byte(fmt.Sprintf("%s%d", input, i)))
		hash := fmt.Sprintf("%x", h)

		if hash[0:5] == "00000" {
			if hash[5] <= '7' {
				fmt.Printf("(%c,%c) ", hash[5], hash[6])
				position := parseInt(string(hash[5]))
				if password[position] == '_' {
					password[position] = rune(hash[6])

					l++
				}
			}
		}
	}
	fmt.Println(string(password))
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
