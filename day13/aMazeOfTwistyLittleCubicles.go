package main

import (
	"fmt"
	"math"
)

var input = 1364
var size = 50

type position struct {
	x int
	y int
}

type maze [][]bool

type Node struct {
	position position
	edges    []*Node
}
type Graph struct {
	nodes map[position]*Node
}

func main() {
	m := buildMaze()
	g := buildGraph(m)

	m.Print()

	distances := dijkstra(g.nodes[position{1, 1}], &g.nodes)
	fmt.Printf("nodeCount: %d\n", len(g.nodes))
	fmt.Printf("answer one: %d\n", distances[position{31, 39}])

	sum := 0

	for _, v := range distances {
		if v <= 50 {
			sum++
		}
	}
	fmt.Printf("answer two: %d\n", sum)
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

func buildMaze() maze {

	m := make([][]bool, size)
	for y := 0; y < size; y++ {
		m[y] = make([]bool, size)
		for x := 0; x < size; x++ {
			m[y][x] = isWall(x, y)
		}
	}
	return m
}

func isWall(x int, y int) bool {
	number := x*x + 3*x + 2*x*y + y + y*y + input
	binary := fmt.Sprintf("%b", number)
	count := 0
	for _, bit := range binary {
		if bit == '1' {
			count++
		}
	}

	return count%2 == 1
}

func buildGraph(m maze) Graph {
	nodes := map[position]*Node{}

	for y := 0; y < size; y++ {
		for x := 0; x < size-1; x++ {
			isWall := m[y][x]
			if isWall {
				continue
			}

			nPos := position{x: x, y: y}
			node, ok := nodes[nPos]
			if !ok {
				node = &Node{position: nPos, edges: []*Node{}}
				nodes[nPos] = node
			}

			rPos := position{x: x + 1, y: y}
			if x < size-1 && !m[y][x+1] {
				nodeToRight, ok := nodes[rPos]
				if !ok {
					nodeToRight = &Node{position: rPos, edges: []*Node{}}

				}
				node.edges = append(node.edges, nodeToRight)
				nodeToRight.edges = append(nodeToRight.edges, node)
				nodes[rPos] = nodeToRight
			}

			dPos := position{x: x, y: y + 1}
			if y < size-1 && !m[y+1][x] {
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

func dijkstra(startNode *Node, nodes *map[position]*Node) map[position]int {
	distances := map[position]int{}

	for k, _ := range *nodes {
		distances[k] = math.MaxInt32
	}
	distances[(*startNode).position] = 0

	for {
		if len(*nodes) == 0 {
			break
		}
		var current position
		min := math.MaxInt32
		for position, _ := range *nodes {
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
		for _, neighbour := range (*nodes)[current].edges {
			if distances[current]+1 < distances[(*neighbour).position] {
				distances[(*neighbour).position] = distances[current] + 1
			}
		}
		delete(*nodes, current)
	}
	return distances
}
