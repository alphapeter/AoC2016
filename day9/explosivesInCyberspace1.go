package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type NodeType int

const (
	VALUE NodeType = 0
	BOT NodeType = 1
	OUTPUT NodeType = 2
)
type Node struct{
	Name string
	OutHigh *Node
	OutLow *Node
	InHigh *Node
	InLow *Node

	Type NodeType

	Value int
}

func main() {
	nodes := make(map[string]Node)
	inputData, _ := ioutil.ReadFile("day10/input.txt")
	input := strings.Split(strings.TrimSpace(string(inputData)), "\n")

	for row := range input{
		command := strings.Fields(row)
		if(command[0] == "bot"){
			//bot
			botNr := parseInt(command[5])
			botNode := getOrCreateNode(nodes, "bot" + botNr, BOT)
			botNode.Type = BOT

			setLow(nodes, botNode, command[5:6])
			setHigh(nodes, botNode, command[10:11])
		}
	}

	for row := range input{
		command := strings.Fields(row)
		if(command[0] == "value"){
			//value
			value := parseInt(command[1])
			valueNode := getOrCreateNode(nodes, "value" + value, VALUE)
			valueNode.Value = value
			//bot
			botNr := parseInt(command[5])
			botNode := getOrCreateNode(nodes, "bot" + botNr, BOT)
			if(botNode.InLow == nil){
				botNode.InLow = valueNode
			}else{
				botNode.InHigh = valueNode
			}
		}
	}

}

func setHigh(nodes map[string]Node, node* Node, command string){
	if(command[0] == bot){

	}
}
func getType(s string){
	switch s {
	case "bot": return BOT

	}
	if(s == "bot"){
		return BOT
	}
	else{

}
}

func getOrCreateNode(bots* map[string]Node, name string, t NodeType) Node{
	bot, ok := bots[name]
	if(!ok){
		bot = Node{Name: "bot" + name}
		bot.Type = t
		bots[name] = bot
	}
	return bot
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
