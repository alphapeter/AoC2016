package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputData, _ := ioutil.ReadFile("day4/input.txt")
	input := string(inputData)

	rows := strings.Split(input, "\n")
	fmt.Println("number of rooms: ", len(rows))

	sum := 0

	for _, row := range rows {
		if row == "" {
			continue
		}
		sectorId, isValid := checkRoom(row)
		if isValid {
			sum = sum + sectorId

		}

	}
	fmt.Println("answer: ", sum)
}

func abs(integer int) int {
	if integer < 0 {
		return -integer
	}
	return integer
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}

func checkRoom(s string) (int, bool) {
	return getSectorId(s), getCheckSum(s) == calculateCheckSum(s)
}

func getCheckSum(s string) string {
	var i int
	for i = 0; s[i] != '['; i++ {
	}
	i++
	return s[i : i+5]
}

func calculateCheckSum(s string) string {
	letters := map[rune]int{}
	var i int

	for i = 0; s[i] > '9' || s[i] == '-'; i++ {
		if s[i] == '-' {
			continue
		} else {
			letters[rune(s[i])]++
		}
	}

	letterFrequency := make(keyValues, len(letters))
	j := 0
	for k, v := range letters {
		letterFrequency[j] = keyValue{Key: rune(k), Value: v}
		j++
	}
	sort.Sort(sort.Reverse(letterFrequency))

	checkSum := make([]rune, 5)

	for l := 0; l < 5; l++ {
		checkSum[l] = letterFrequency[l].Key
	}
	return string(checkSum)
}

func getSectorId(s string) int {
	var i int
	var j int

	for i = 0; s[i] > '9' || s[i] == '-'; i++ {
	}

	for j = i; s[j] != '['; j++ {

	}
	return parseInt(s[i:j])
}

type keyValue struct {
	Key   rune
	Value int
}

type keyValues []keyValue

func (p keyValues) Len() int { return len(p) }
func (p keyValues) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}
	return p[i].Value < p[j].Value
}
func (p keyValues) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
