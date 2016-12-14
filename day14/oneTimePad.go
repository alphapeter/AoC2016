package main

import (
	"crypto/md5"
	"fmt"
	"regexp"
)

type hashFn func(string, int) string

var hashCache1 map[int]string = map[int]string{}
var hashCache2 map[int]string = map[int]string{}
var regex3 *regexp.Regexp = regexp.MustCompile(`(0{3}|1{3}|2{3}|3{3}|4{3}|5{3}|6{3}|7{3}|8{3}|9{3}|a{3}|b{3}|c{3}|d{3}|e{3}|f{3})`)

//var salt = "abc"
var salt = "cuanljph"

func main() {

	answer1 := count(getHash1)
	answer2 := count(getHash2)

	fmt.Printf("answer1: %d\n answer2: %d", answer1, answer2)

}

func count(getHash hashFn) int {
	l := 0
	for i := 0; l < 64; i++ {
		hash := getHash(salt, i)
		if regex3.MatchString(hash) {
			character := rune(regex3.FindString(hash)[0])
			regex5 := regexp.MustCompile(fmt.Sprintf("(%c{5})", character))
			for j := 1; j <= 1000; j++ {
				_hash := getHash(salt, i+j)
				if regex5.MatchString(_hash) {
					//fmt.Printf("%d: found %s at index: %d\n", l+1, _hash, i)
					l++
					if l == 64 {
						return i
					}
					break
				}
			}
		}
	}
	return 0
}

func getHash1(salt string, i int) string {
	hash, ok := hashCache1[i]
	if ok {
		return hash
	}
	h := md5.Sum([]byte(fmt.Sprintf("%s%d", salt, i)))
	hash = fmt.Sprintf("%x", h)
	hashCache1[i] = hash
	return hash
}

func getHash2(salt string, i int) string {
	hash, ok := hashCache2[i]
	if ok {
		return hash
	}

	hash = fmt.Sprintf("%s%d", salt, i)
	for i := 0; i < 2017; i++ {
		hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	}
	hashCache2[i] = hash
	return hash
}
