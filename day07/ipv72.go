package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inputData, _ := ioutil.ReadFile("day07/input.txt")
	input := string(inputData)

	rows := strings.Split(input, "\n")
	fmt.Println("number of addresses: ", len(rows))

	sum := 0

	for _, row := range rows {
		if row == "" {
			continue
		}
		for i := 0; i < len(row)-2; i++ {
			if row[i] == '[' {
				i = strings.Index(row[i:], "]") + i
			}
			pattern, isMatch := checkAba(row, i)
			if isMatch && checkHyphens(row, pattern) {
				sum++
				break

			}
		}

	}
	fmt.Println("answer: ", sum)
}

func checkAba(row string, i int) (string, bool) {
	return row[i : i+3], row[i] == row[i+2] && row[i] != row[i+1] && row[i+1] != '['
}

func checkHyphens(row string, pattern string) bool {
	for _, hyphen := range strings.Split(row, "[") {
		j := strings.Index(hyphen, "]")
		for i := 0; i < j-2; i++ {
			if hyphen[i] == pattern[1] && hyphen[i+1] == pattern[0] && hyphen[i+2] == pattern[1] {
				return true
			}
		}
	}
	return false
}
