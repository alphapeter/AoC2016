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
	t1 := triangle{a: 0, b: 0, c: 0}
	t2 := triangle{a: 0, b: 0, c: 0}
	t3 := triangle{a: 0, b: 0, c: 0}

	for _, row := range rows {
		sides := strings.Fields(row)
		if len(sides) != 3 {
			continue
		}
		s1 := parsInt(sides[0])
		s2 := parsInt(sides[1])
		s3 := parsInt(sides[2])
		if t1.a == 0 {
			t1.a = s1
			t2.a = s2
			t3.a = s3
		} else if t1.b == 0 {
			t1.b = s1
			t2.b = s2
			t3.b = s3
		} else if t1.c == 0 {
			t1.c = s1
			t2.c = s2
			t3.c = s3

			if isValid(&t1) {
				sum++
			}
			if isValid(&t2) {
				sum++
			}
			if isValid(&t3) {
				sum++
			}

			reset(&t1)
			reset(&t2)
			reset(&t3)
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

func isValid(t *triangle) bool {
	return t.a+t.b > t.c && abs(t.a-t.b) < t.c
}

func reset(t *triangle) {
	t.a = 0
	t.b = 0
	t.c = 0
}
