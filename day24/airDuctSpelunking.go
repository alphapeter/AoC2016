package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type maze [][]bool

type Node struct {
	position   position
	edges      []*Node
	number     int
	isLocation bool
}
type Graph struct {
	nodes map[position]*Node
}

func main() {

	g := buildGraph()

	locations := [8]position{}

	for _, node := range g.nodes {
		if node.isLocation {
			locations[node.number] = node.position
		}
	}

	distances := [8][8]int{}
	for i, location := range locations {
		distances[i] = [8]int{}

		distancesToOtherLocations := dijkstra(g.nodes[location], copyNodes(g.nodes))
		for j, other := range locations {
			distances[i][j] = distancesToOtherLocations[other]
		}
	}

	a := []int{1, 2, 3, 4, 5, 6, 7}
	permutations := createPermutations([]int{}, a)

	distance := 1000000
	for _, p := range permutations {
		path := append([]int{0}, p...)
		d := getLength(path, distances)
		if d < distance {
			distance = d
		}
	}

	fmt.Printf("answer one: %d\n", distance)

	distance = 1000000
	for _, p := range permutations {
		path := append([]int{0}, p...)
		path = append(path, 0)
		d := getLength(path, distances)
		if d < distance {
			distance = d
		}
	}
	fmt.Printf("answer two: %d\n", distance)

}

func getLength(path []int, distances [8][8]int) int {

	distance := 0
	for i := 1; i < len(path); i++ {
		from := path[i-1]
		to := path[i]

		distance += distances[from][to]
	}
	return distance
}

func createPermutations(numbers []int, rest []int) [][]int {
	if len(rest) == 0 {
		return [][]int{numbers}
	}

	permutations := [][]int{}
	for i := 0; i < len(rest); i++ {
		_rest := make([]int, len(rest))
		_numbers := make([]int, len(numbers))
		copy(_rest, rest)
		copy(_numbers, numbers)

		permutation := append(_numbers, _rest[i])
		newRest := append(_rest[:i], _rest[i+1:]...)
		newPermutations := createPermutations(permutation, newRest)
		permutations = append(permutations, newPermutations...)
	}
	return permutations
}

func (m *maze) Print() {
	fmt.Print("  ")
	for x := 0; x < len((*m)[0]); x++ {
		fmt.Printf("%2d", x)
	}
	fmt.Println()
	for y := 0; y < len(*m); y++ {
		fmt.Printf("%2d", y)
		for x := 0; x < len((*m)[y]); x++ {
			if (*m)[y][x] {
				fmt.Print(" #")
			} else {
				fmt.Print(" .")
			}
		}
		fmt.Print("\n")
	}
}

func buildGraph() Graph {
	nodes := map[position]*Node{}

	inputData, _ := ioutil.ReadFile("day24/input.txt")
	input := string(inputData)

	rows := strings.Split(input, "\n")

	for y, row := range rows {
		for x, n := range row {
			isWall := n == '#'
			if isWall {
				continue
			}

			nPos := position{x: x, y: y}
			node, ok := nodes[nPos]
			if !ok {
				node = &Node{position: nPos, edges: []*Node{}}
				nodes[nPos] = node
			}

			if n != '.' {
				node.number = parseInt(string(n))
				node.isLocation = true
			}

			rPos := position{x: x + 1, y: y}
			if x < len(row)-1 && rows[y][x+1] != '#' {
				nodeToRight, ok := nodes[rPos]
				if !ok {
					nodeToRight = &Node{position: rPos, edges: []*Node{}}

				}
				node.edges = append(node.edges, nodeToRight)
				nodeToRight.edges = append(nodeToRight.edges, node)
				nodes[rPos] = nodeToRight
			}

			dPos := position{x: x, y: y + 1}
			if y < len(rows)-1 && rows[y+1][x] != '#' {
				nodeBeneath, ok := nodes[dPos]
				if !ok {
					nodeBeneath = &Node{position: dPos, edges: []*Node{}}

				}
				node.edges = append(node.edges, nodeBeneath)
				nodeBeneath.edges = append(nodeBeneath.edges, node)
				nodes[dPos] = nodeBeneath
			}
		}

	}
	return Graph{nodes: nodes}
}

func dijkstra(startNode *Node, nodes map[position]*Node) map[position]int {
	distances := map[position]int{}

	for k, _ := range nodes {
		distances[k] = math.MaxInt32
	}
	distances[(*startNode).position] = 0

	for {
		if len(nodes) == 0 {
			break
		}
		var current position
		min := math.MaxInt32
		for position, _ := range nodes {
			dist := distances[position]
			if dist < min {
				min = distances[position]
				current = position
			}
		}

		//no more reachable nodes
		if min == math.MaxInt32 {
			break
		}
		//get minimum distances
		for _, neighbour := range nodes[current].edges {
			if distances[current]+1 < distances[(*neighbour).position] {
				distances[(*neighbour).position] = distances[current] + 1
			}
		}
		delete(nodes, current)
	}
	return distances
}

func copyNodes(nodes map[position]*Node) map[position]*Node {
	newMap := map[position]*Node{}
	for k, v := range nodes {
		newMap[k] = v
	}
	return newMap
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
