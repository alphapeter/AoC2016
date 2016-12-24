package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"sort"
)

type Range struct {
	Min uint32
	Max uint32
}

type Ranges []Range

func (p Ranges) Len() int { return len(p) }
func (p Ranges) Less(i, j int) bool {
	return p[i].Min < p[j].Min
}
func (p Ranges) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func createRange(row []string) Range {
	return Range{Min: parseInt(row[0]), Max: parseInt(row[1])}
}

func main() {
	inputData, _ := ioutil.ReadFile("day20/input.txt")
	input := strings.Split(string(inputData), "\n")


	rs := []Range{}
	for _, row := range input {
		if row == "" {
			continue
		}
		rs = append(rs, createRange(strings.Split(row, "-")))

	}
	ranges := Ranges(rs)
	sort.Sort(ranges)

	solve1(ranges)
	solve2(ranges)
}

func solve1(ranges Ranges){
	var low uint32 =  0
	for _, r := range ranges{
		if low + 1 >= r.Min {
			low = r.Max
		}else{
			fmt.Printf("solution1: %d\n", low +1)
			break
		}
	}
}
func solve2(ranges Ranges) {
	var sum uint32 = 0
	var low uint32 = 0

	for _, r := range ranges{
		if low+ 1 >= r.Min {

		}else if r.Min > low {
			sum = sum + (r.Min - low -1)
		}
		if r.Max > low {
			low = r.Max
		}
	}

	sum = sum + (4294967295 - low)

	fmt.Printf("solution2: %d", sum)
}

func parseInt(s string) uint32 {
	a, _ := strconv.ParseInt(s, 10, 64)
	return uint32(a)
}
