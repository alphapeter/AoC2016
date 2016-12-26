package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"regexp"
)

type Password []rune

func main() {
	inputData, _ := ioutil.ReadFile("day22/input.txt")
	input := strings.Split(string(inputData), "\n")
	nodes := parseNodes(input)
	solve1(nodes)
}

type Node struct{
	X int
	Y int
	Size int
	Used int
	Avail int
	Percentage int
}

func solve1(nodes []Node) {
	count := 0
	for _, a := range nodes{
		for _, b := range nodes{
			if a.Used > 0 && a.Used <= b.Avail && !(a.X == b.X && a.Y == b.Y) {
				count++
			}
		}
	}

	fmt.Printf("answer1: %d\n", count)
}


func parseNodes(input []string)[]Node{
	nodes := []Node{}
	for _, row := range input{
		if !strings.HasPrefix(row, "/dev") {
			continue
		}
		nodes = append(nodes, parseNode(row))
	}
	return nodes
}

var xPositionRegex = regexp.MustCompile(`x(\d+)`)
var yPositionRegex = regexp.MustCompile(`y(\d+)`)
var sizeRegex = regexp.MustCompile(` (\d+)`)

func parseNode(row string) Node{
	x := parseInt(xPositionRegex.FindStringSubmatch(row)[1])
	y := parseInt(yPositionRegex.FindStringSubmatch(row)[1])
	sizes := sizeRegex.FindAllStringSubmatch(row, -1)
	size := parseInt(sizes[0][1])
	used := parseInt(sizes[1][1])
	avail := parseInt(sizes[2][1])
	percentage := parseInt(sizes[3][1])

	return Node{X: x, Y: y, Size: size, Avail: avail, Used: used, Percentage: percentage}
}



func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
