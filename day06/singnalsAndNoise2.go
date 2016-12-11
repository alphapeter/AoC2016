package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inputData, _ := ioutil.ReadFile("day6/input.txt")
	input := string(inputData)
	rows := strings.Split(input, "\n")

	letters := make(map[int]map[rune]int)

	for i := 0; i < 8; i++ {
		letters[i] = map[rune]int{}
	}

	for _, row := range rows {
		for i := 0; i < len(row); i++ {
			letter := rune(row[i])
			letters[i][letter]++
		}
	}

	for i := 0; i < len(letters); i++ {
		var l rune
		count := 500
		for k, v := range letters[i] {
			if v < count {
				count = v
				l = k
			}
		}
		fmt.Printf("%c", rune(l))
	}
}
