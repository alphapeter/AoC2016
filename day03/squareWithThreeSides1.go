package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type triangle struct {
	a int
	b int
	c int
}

func main() {
	inputData, _ := ioutil.ReadFile("day03/input.txt")
	input := string(inputData)

	rows := strings.Split(input, "\n")
	fmt.Println("number of triangles: ", len(rows))
	sum := 0

	for _, row := range rows {
		sides := strings.Fields(row)
		if len(sides) != 3 {
			continue
		}

		t := triangle{
			a: parsInt(sides[0]),
			b: parsInt(sides[1]),
			c: parsInt(sides[2])}
		if isValid(t) {
			sum++
		}

	}
	fmt.Println("real triangles: ", sum)
}

func abs(integer int) int {
	if integer < 0 {
		return -integer
	}
	return integer
}

func parsInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}

func isValid(t triangle) bool {
	return t.a+t.b > t.c && abs(t.a-t.b) < t.c
}
