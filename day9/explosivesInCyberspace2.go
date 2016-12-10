package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputData, _ := ioutil.ReadFile("day9/input.txt")
	input := string(inputData)

	data := strings.TrimSpace(input)

	//test cases
	//data := "(3x3)XYZ" //9
	//data := "X(8x2)(3x3)ABCY" //20
	//data := "(27x12)(20x12)(13x14)(7x10)(1x12)A" //241920
	//data := "(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN" //445
	//11451628995
	length := decompress(data)

	fmt.Println("answer:", length)
}

func decompress(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			length, multiplier, offset := parseMarker(s[i:])
			j := i + offset + 1
			return i + decompress(appendMultiple(s[j:j+length], multiplier)) + decompress(s[j+length:])
		}
	}
	return len(s)
}

func appendMultiple(s string, multiplier int) string {
	if multiplier == 0 {
		return ""
	}
	return s + appendMultiple(s, multiplier-1)
}

func parseMarker(s string) (int, int, int) {

	markerEnd := strings.Index(s, ")")
	marker := s[1:markerEnd]
	numbers := strings.Split(marker, "x")
	length := parseInt(numbers[0])
	multiplier := parseInt(numbers[1])

	return length, multiplier, markerEnd

}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
