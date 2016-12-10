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
	//data := "ADVENT"
	//data := "A(1x5)BC"
	//data := "(3x3)XYZ"
	//data := "(6x1)(1x3)A"
	//data := "X(8x2)(3x3)ABCY"

	v := "s" + "sf"
	fmt.Print(v)

	decompressed := decompress(data)

	fmt.Println("answer:", len(decompressed))
}

func decompress(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			length, multiplier, offset := parseMarker(s[i:])
			j := i + offset + 1
			return s[0:i] + appendMultiple(s[j:j+length], multiplier) + decompress(s[j+length:])
		}
	}
	return s
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
