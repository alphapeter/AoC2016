package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inputData, _ := ioutil.ReadFile("day7/input.txt")
	input := string(inputData)

	rows := strings.Split(input, "\n")
	fmt.Println("number of addresses: ", len(rows))

	sum := 0

	for _, row := range rows {
		if row == "" {
			continue
		}
		for i := 0; i < len(row)-3; i++ {
			if row[i] == '[' {
				i = strings.Index(row[i:], "]") + i
			}
			if checkAbba(row, i) && !checkHyphens(row) {
				sum++
				break
			}
		}

	}
	fmt.Println("answer: ", sum)
}

func checkAbba(row string, i int) bool {
	return row[i] == row[i+3] && row[i+1] == row[i+2] && row[i] != row[i+1]
}

func checkHyphens(row string) bool {
	for _, hyphen := range strings.Split(row, "[") {
		j := strings.Index(hyphen, "]")
		for i := 0; i < j-3; i++ {
			if checkAbba(hyphen, i) {
				return true
			}
		}
	}
	return false
}
