package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strings"
)

//(g1,m1),(g2,m2),(g3,m3),(g4,m5),level

type State [9]int

var G1, M1, G2, M2, G3, M3, G4, M4, E = 0, 1, 2, 3, 4, 5, 6, 7, 8

var moves *[][9]int

type Node struct {
	state *State
	edges []*Node
}

func (s *State) isValid() bool {
	return (s[G1] >= s[M1] || s[G1] == 0) &&
		(s[G2] >= s[M2] || s[G2] == 0) &&
		(s[G3] >= s[M3] || s[G3] == 0) &&
		(s[G4] >= s[M4] || s[G4] == 0) &&
		s[G1] >= 0 && s[M1] >= 0 &&
		s[G2] >= 0 && s[M2] >= 0 &&
		s[G3] >= 0 && s[M3] >= 0 &&
		s[G4] >= 0 && s[M4] >= 0 &&
		s[E] >= 0 && s[E] < 4
}

func (s *State) Print() {
	fmt.Printf("\n")
	fmt.Printf("|level|Elevator|Generators|Microchips|\n")
	fmt.Printf("| 4   |   %s    |    %d     |    %d     |\n", printElevator(3, s[E]), s[G4], s[M4])
	fmt.Printf("| 3   |   %s    |    %d     |    %d     |\n", printElevator(2, s[E]), s[G3], s[M3])
	fmt.Printf("| 2   |   %s    |    %d     |    %d     |\n", printElevator(1, s[E]), s[G2], s[M2])
	fmt.Printf("| 1   |   %s    |    %d     |    %d     |\n", printElevator(0, s[E]), s[G1], s[M1])
}

func printElevator(i int, e int) string {
	if i == e {
		return "E"
	}
	return " "
}

func main() {
	microChipRegex := regexp.MustCompile("microchip")
	generatorRegex := regexp.MustCompile("generator")

	inputData, _ := ioutil.ReadFile("day11/input.txt")
	input := strings.Split(strings.TrimSpace(string(inputData)), "\n")

	state := State([9]int{0, 0, 0, 0, 0, 0, 0, 0, 0})

	for i, row := range input {
		chips := len(microChipRegex.FindAllStringSubmatch(row, -1))
		state[i*2+1] = chips
		generators := len(generatorRegex.FindAllStringSubmatch(row, -1))
		state[i*2] = generators
	}

	for i := 0; i < 2; i++ {
		extra := i * 2

		state[0], state[1] = state[0]+extra, state[1]+extra

		startNode := Node{state: &state, edges: []*Node{}}

		nodes := map[State]*Node{}
		moves = generateMoves()
		fmt.Printf("number of states:%d", generateGraph(&startNode, &nodes))

		state.Print()

		distances := dijkstra(&startNode, &nodes)

		fmt.Printf("Answer part %d: %d\n", i+1, distances[State([9]int{0, 0, 0, 0, 0, 0, 5 + extra, 5 + extra, 3})])
	}

}

func dijkstra(startNode *Node, nodes *map[State]*Node) map[State]int {
	distances := map[State]int{}

	for k, _ := range *nodes {
		distances[k] = math.MaxInt32
	}
	distances[*startNode.state] = 0

	for {
		if len(*nodes) == 0 {
			break
		}
		var current State
		min := math.MaxInt32
		for state, _ := range *nodes {

			dist := distances[state]
			if dist < min {
				min = distances[state]
				current = state
			}
		}

		//get minimum distances
		for _, node := range (*nodes)[current].edges {
			if distances[current]+1 < distances[*node.state] {
				distances[*node.state] = distances[current] + 1
			}
		}
		delete(*nodes, current)

	}

	return distances
}

func generateGraph(n *Node, visited *map[State]*Node) int {
	numberOfStates := 0
	for _, move := range *moves {
		m := State(move)
		newState := add(n.state, &m)

		if !newState.isValid() {
			continue
		}

		node, ok := (*visited)[*newState]
		if !ok {
			node = &Node{state: newState, edges: []*Node{}}
			(*visited)[*newState] = node
			numberOfStates++
			numberOfStates = numberOfStates + generateGraph(node, visited)

		}
		n.edges = append(n.edges, node)
	}
	return numberOfStates
}

func add(s1 *State, s2 *State) *State {
	newState := State([9]int{s1[0] + s2[0], s1[1] + s2[1], s1[2] + s2[2], s1[3] + s2[3], s1[4] + s2[4], s1[5] + s2[5], s1[6] + s2[6], s1[7] + s2[7], s1[8] + s2[8]})
	return &newState
}

func generateMoves() *[][9]int {
	m1 := [][]int{{0, -1, 0, 1}, {-1, 0, 1, 0}, {-2, 0, 2, 0}, {0, -2, 0, 2}}

	moves := [][9]int{}
	for i := 0; i < 3; i++ {
		for _, m := range m1 {
			empty := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			empty[i*2], empty[i*2+1], empty[i*2+2], empty[i*2+3], empty[E] = m[0], m[1], m[2], m[3], 1
			moves = append(moves, empty)
		}
		for _, m := range m1 {
			empty := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
			empty[i*2], empty[i*2+1], empty[i*2+2], empty[i*2+3], empty[E] = m[0]*-1, m[1]*-1, m[2]*-1, m[3]*-1, -1
			moves = append(moves, empty)
		}
	}
	return &moves
}
