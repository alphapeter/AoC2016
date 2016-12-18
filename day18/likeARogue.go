package main

import (
	"fmt"
	"io/ioutil"
)

type Row []bool

func (r *Row) getSafeCount() int {
	sum := 0
	for i := 1; i < len(*r)-1; i++ {
		if !(*r)[i] {
			sum++
		}
	}
	return sum
}

func main() {
	inputData, _ := ioutil.ReadFile("day18/input.txt")
	input := string(inputData)

	fmt.Printf("Answer 1: %d\n", solve(input, 40))
	fmt.Printf("Answer 2: %d\n", solve(input, 400000))
}

func solve(input string, size int) int {
	firstRow := Row(make([]bool, len(input)+2))

	firstRow[0], firstRow[len(firstRow)-1] = false, false // initiate imaginary tiles

	for i := 0; i < len(input); i++ {
		if input[i] == '^' {
			firstRow[i+1] = true
		} else {
			firstRow[i+1] = false
		}

	}

	rows := make([]Row, size)
	rows[0] = firstRow
	width := len(firstRow)
	for i := 1; i < size; i++ {
		r := make([]bool, width)
		r[0], r[width-1] = false, false
		for j := 1; j < width-1; j++ {
			r[j] = rows[i-1][j-1] && rows[i-1][j] && !rows[i-1][j+1] ||
				!rows[i-1][j-1] && rows[i-1][j] && rows[i-1][j+1] ||
				rows[i-1][j-1] && !rows[i-1][j] && !rows[i-1][j+1] ||
				!rows[i-1][j-1] && !rows[i-1][j] && rows[i-1][j+1]
		}
		rows[i] = Row(r)
	}

	sum := 0
	for i := 0; i < size; i++ {
		sum = sum + rows[i].getSafeCount()
	}
	return sum
}
