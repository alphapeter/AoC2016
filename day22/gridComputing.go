package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	X          int
	Y          int
	Size       int
	Used       int
	Avail      int
	Percentage int
}

func main() {
	inputData, _ := ioutil.ReadFile("day22/input.txt")
	input := strings.Split(string(inputData), "\n")
	nodes := parseNodes(input)
	printGrid(nodes)
	solve1(nodes)
	solve2(nodes)
}
func printGrid(nodes map[int]map[int]Node) {
	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[y]); x++ {
			n := nodes[y][x]
			n.print()
		}
		fmt.Print("\n")
	}
}

func solve1(nodes map[int]map[int]Node) {
	count := 0
	for _, ay := range nodes {
		for _, a := range ay {
			for _, by := range nodes {
				for _, b := range by {
					if a.Used > 0 && a.Used <= b.Avail && !(a.X == b.X && a.Y == b.Y) {
						count++
					}
				}
			}
		}
	}

	fmt.Printf("answer1: %d\n", count)
}

func solve2(nodes map[int]map[int]Node) {
	emptyNode := findEmptyNode(nodes)
	stepsMoveToLeftOfTopRight := StepsForMoveToTopRight(nodes, emptyNode) - 1
	width := len(nodes[0])
	stepsForMoveToTopLeft := width - 1 + (width-2)*4
	fmt.Printf("answer2: %d", stepsMoveToLeftOfTopRight+stepsForMoveToTopLeft)
}

// Assumes that the passage is blocked horizontally with no passage to the right
func StepsForMoveToTopRight(nodes map[int]map[int]Node, node Node) int {
	steps := 0
	for node.Y > 0 {
		n := nodes[node.Y-1][node.X]
		if n.Size > 100 {
			n = nodes[node.Y][node.X-1]
		}
		node = n
		steps++
	}
	var width = len(nodes[0])
	return steps + width - (node.X + 1)
}
func findEmptyNode(nodes map[int]map[int]Node) Node {
	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[y]); x++ {
			node := nodes[y][x]
			if node.Used == 0 {
				return node
			}
		}
	}
	panic("no empty node!")
}

func parseNodes(input []string) map[int]map[int]Node {
	nodes := map[int]map[int]Node{}
	for _, row := range input {
		if !strings.HasPrefix(row, "/dev") {
			continue
		}
		node := parseNode(row)
		_, ok := nodes[node.Y]
		if !ok {
			nodes[node.Y] = map[int]Node{}
		}
		nodes[node.Y][node.X] = node
	}
	return nodes
}

var xPositionRegex = regexp.MustCompile(`x(\d+)`)
var yPositionRegex = regexp.MustCompile(`y(\d+)`)
var sizeRegex = regexp.MustCompile(` (\d+)`)

func parseNode(row string) Node {
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
func (node *Node) print() {
	if node.Size < 70 {
		fmt.Print(":XX/XX:")
	} else if node.Avail >= 70 {
		fmt.Printf("*%2d/%2d*", node.Used, node.Size)
	} else if node.Used >= 100 {
		fmt.Print(":XX/XX:")
	} else {
		fmt.Printf("(%2d/%2d)", node.Used, node.Size)
	}
}
