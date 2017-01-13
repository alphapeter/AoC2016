package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputData, _ := ioutil.ReadFile("day23/input.txt")
	input := string(inputData)

	rows := strings.Split(input, "\n")
	fmt.Printf("answer 1: %d\n", run(getCommands(rows), map[string]int{"a": 7}))
	fmt.Printf("answer 2: %d\n", run(getCommands(rows), map[string]int{"a": 12}))
}

func getCommands(rows []string) [][]string {
	commands := [][]string{}
	for _, row := range rows {
		if row != "" {
			commands = append(commands, strings.Fields(row))
		}
	}
	return commands
}

func run(commands [][]string, registers map[string]int) int {
	for pPointer := 0; pPointer < len(commands); pPointer++ {
		command := commands[pPointer]
		if len(command) == 0 {
			continue
		}
		switch command[0] {
		case "cpy":
			source := command[1]
			destination := command[2]
			if rune(destination[0]) <= '9' || rune(destination[0]) == '-' {
				continue
			}
			var val int
			if rune(source[0]) >= 'a' {
				val = registers[source]
			} else {
				val = parseInt(source)
			}
			registers[destination] = val
			break

		case "inc":
			registers[command[1]] = registers[command[1]] + 1
			break
		case "dec":
			registers[command[1]] = registers[command[1]] - 1
			break
		case "jnz":
			var val int
			if rune(command[1][0]) < 'a' { // number
				val = parseInt(command[1])
				if val != 0 {
					pPointer = pPointer + getOffset(command, registers) - 1
				}
			} else {
				val = registers[command[1]]
				if val != 0 {
					if incBy(pPointer, command, commands) { //optimize increase by (dec inc jnz or inc dec jnz)
						incByCmd := commands[pPointer-1]
						if commands[pPointer-1][0] == "dec" {
							incByCmd = commands[pPointer-2]
						}
						registers[incByCmd[1]] += registers[command[1]]
						registers[command[1]] = 0

					} else if increaseByMultiplication(pPointer, command, commands) { //optimize multiplication
						inc := commands[pPointer-4][1]
						factor1 := command[1]
						factor2 := commands[pPointer-5][1]

						registers[inc] += registers[factor1] * registers[factor2]
						registers[command[1]] = 0
					} else {
						pPointer = pPointer + getOffset(command, registers) - 1
					}
				}
			}
			break
		case "tgl":
			insPointer := pPointer + registers[command[1]]
			if insPointer >= len(commands) || insPointer <= 0 {
				continue
			}

			tglRow := commands[insPointer]
			tglCommand := tglRow[0]
			tglArguments := tglRow[1:]

			if len(tglArguments) == 1 {
				if tglCommand == "inc" {
					commands[insPointer][0] = "dec"

				} else {
					commands[insPointer][0] = "inc"
				}
			} else { //binary
				if tglCommand == "jnz" {
					commands[insPointer][0] = "cpy"
				} else {
					commands[insPointer][0] = "jnz"
				}
			}
		default:
			panic(fmt.Sprintf("command not implemented %s", command))
		}
	}
	return registers["a"]
}
func increaseByMultiplication(pPointer int, command []string, commands [][]string) bool {
	return parseInt(command[2]) == -5 && rune(commands[pPointer-5][1][0]) >= 'a'
}
func incBy(pPointer int, command []string, commands [][]string) bool {
	return parseInt(command[2]) == -2 && ((commands[pPointer-1][0] == "dec" && commands[pPointer-2][0] == "inc") || (commands[pPointer-1][0] == "inc" && commands[pPointer-2][0] == "dec"))
}

func getOffset(command []string, registers map[string]int) int {
	var offset int
	if rune(command[2][0]) >= 'a' {
		offset = registers[command[2]]
	} else {
		offset = parseInt(command[2])
	}
	return offset
}

func parseInt(s string) int {
	a, _ := strconv.ParseInt(s, 10, 32)
	return int(a)
}
