package main

import (
	"fmt"
)

func main() {
	input := "11110010111001001"
	solve(input, 272, 1)
	solve(input, 35651584, 2)
}

func solve(input string, length int, problemNo int){
	result := input
	for len(result) < length{
		result = result + "0" + Invert(Reverse(result))
	}
	result = result[:length]


	fmt.Printf("Answer%d: %s\n", problemNo, getChecksum(result))
}

func getChecksum(s string) string{
	if(len(s)%2 == 1){
		return s
	}
	chkSum := make([]rune, len(s)/2)
	for j, i:=0,0; i< len(s); i = i+2{
		if(s[i] == s[i+1]){
			chkSum[j] = '1'
		}else{
			chkSum[j] = '0'
		}
		j++
	}
	return getChecksum(string(chkSum))
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Invert(s string) string {
	runes := []rune(s)
	for i := 0; i< len(s); i++{
		if(runes[i] == '0'){
			runes[i] = '1'
		}else {
			runes[i] = '0'

		}
	}
	return string(runes)
}
