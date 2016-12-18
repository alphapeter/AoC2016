package main

import (
	"crypto/md5"
	"fmt"
	"math"
)

func main() {
	input := "edjrjqaa"
	_, answer1 := getPath(0, 0, input)
	fmt.Printf("solution1: %s\n", answer1)
	_, answer2 := getLongestPath(0, 0, input)
	fmt.Printf("solution2: %d\n", answer2)
}

func getPath(x int, y int, passPhrase string) (bool, string) {
	if x < 0 || y < 0 || x > 3 || y > 3 {
		return false, ""
	}

	if x == 3 && y == 3 {
		return true, ""
	}

	path := ""
	directions := fmt.Sprintf("%x", md5.Sum([]byte(passPhrase)))[:4]
	if isOpen(rune(directions[0])) {
		isOk, pathU := getPath(x, y-1, passPhrase+"U")
		if isOk && (path == "" || len(pathU) <= len(path)) {
			path = "U" + pathU
		}
	}
	if isOpen(rune(directions[1])) {
		isOk, pathD := getPath(x, y+1, passPhrase+"D")
		if isOk && (path == "" || len(pathD) <= len(path)) {
			path = "D" + pathD
		}
	}
	if isOpen(rune(directions[2])) {
		isOk, pathL := getPath(x-1, y, passPhrase+"L")
		if isOk && (path == "" || len(pathL) <= len(path)) {
			path = "L" + pathL
		}
	}
	if isOpen(rune(directions[3])) {
		isOk, pathR := getPath(x+1, y, passPhrase+"R")
		if isOk && (path == "" || len(pathR) <= len(path)) {
			path = "R" + pathR
		}
	}
	return path != "", path
}

func getLongestPath(x int, y int, passPhrase string) (bool, int) {
	if x < 0 || y < 0 || x > 3 || y > 3 {
		return false, 0
	}

	if x == 3 && y == 3 {
		return true, 0
	}

	path := math.MaxInt32
	directions := fmt.Sprintf("%x", md5.Sum([]byte(passPhrase)))[:4]
	if isOpen(rune(directions[0])) {
		isOk, pathU := getLongestPath(x, y-1, passPhrase+"U")
		if isOk && (path == math.MaxInt32) {
			path = 1 + pathU
		}
	}
	if isOpen(rune(directions[1])) {
		isOk, pathD := getLongestPath(x, y+1, passPhrase+"D")
		if isOk && (path == math.MaxInt32 || pathD >= path) {
			path = 1 + pathD
		}
	}
	if isOpen(rune(directions[2])) {
		isOk, pathL := getLongestPath(x-1, y, passPhrase+"L")
		if isOk && (path == math.MaxInt32 || pathL >= path) {
			path = 1 + pathL
		}
	}
	if isOpen(rune(directions[3])) {
		isOk, pathR := getLongestPath(x+1, y, passPhrase+"R")
		if isOk && (path == math.MaxInt32 || pathR >= path) {
			path = 1 + pathR
		}
	}
	return path != math.MaxInt32, path
}

func isOpen(letter rune) bool {
	return letter >= 'b'
}
