package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type NodeType int

const (
	BOT    NodeType = 0
	OUTPUT NodeType = 1
)

type Node struct {
	Name    string
	HighOut *Node
	LowOut  *Node

	Type NodeType

	HighValue int
	LowValue  int
}

func main() {
	nodes := make(map[string]*Node)
	inputData, _ := ioutil.ReadFile("day10/input.txt")
	input := strings.Split(strings.TrimSpace(string(inputData)), "\n")

	for _, row := range input {
		command := strings.Fields(row)
		if command[0] == "bot" {
			botNr := command[1]
			botNode := getOrCreateNode(&nodes, "bot"+botNr, BOT)
			botNode.Type = BOT
			setLow(&nodes, botNode, command[5:7])
			setHigh(&nodes, botNode, command[10:12])
		}
	}

	var values [][]string
	for _, row := range input {
		command := strings.Fields(row)
		if command[0] == "value" {
			values = append(values, command)
		}
	}
	sort.Sort(ValueSort(values))
	for _, command := range values {
		botNr := command[5]
		botNode, _ := nodes["bot"+botNr]
		value := parseInt(command[1])
		setValue(botNode, value)
	}

	for _, node := range nodes {
		if node.Type == BOT && node.LowValue == 17 && node.HighValue == 61 {
			fmt.Printf("Answer part one: %s, %d, %d\n", node.Name, node.LowValue, node.HighValue)
		}
	}
	product := 1
	for _, n := range nodes {
		if n.Type == OUTPUT && (n.Name == "output0" || n.Name == "output1" || n.Name == "output2") {
			product = product * n.LowValue
		}
	}

	fmt.Println("Answer part two: ", product)
}

func setValue(node *Node, value int) {
	if node == nil {
		return
	}
	if node.LowValue == 0 {
		node.LowValue = value
		setValue(node.LowOut, value)

	} else {
		node.HighValue = value
		setValue(node.HighOut, value)
	}
}

func setLow(nodes *map[string]*Node, node *Node, command []string) {
	childNode := getOrCreateChild(nodes, command)
	node.LowOut = childNode
}

func setHigh(nodes *map[string]*Node, node *Node, command []string) {
	childNode := getOrCreateChild(nodes, command)
	node.HighOut = childNode
}

func getOrCreateChild(nodes *map[string]*Node, command []string) *Node {
	t := getType(command[0])
	return getOrCreateNode(nodes, command[0]+command[1], t)
}

func getType(s string) NodeType {
	switch s {
	case "bot":
		return BOT
	case "output":
		return OUTPUT
	default:
		panic("unhandled switch case")
	}
}

func getOrCreateNode(bots *map[string]*Node, name string, t NodeType) *Node {
	bot, ok := (*bots)[name]
	if !ok {
		bot = &Node{Name: name}
		bot.Type = t
		(*bots)[name] = bot
	}
	return bot
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}

// sorting
type ValueSort [][]string

func (p ValueSort) Len() int { return len(p) }
func (p ValueSort) Less(i, j int) bool {
	return parseInt(p[i][1]) < parseInt(p[j][1])
}
func (p ValueSort) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
