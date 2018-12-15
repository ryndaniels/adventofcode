package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	fmt.Println("advent of code day 5")

	dat, _ := ioutil.ReadFile("./day5.input")
	polymer := string(dat)
	polymer = strings.TrimSuffix(polymer, "\n")

	reactedPolymer := react(polymer)
	fmt.Printf("reacted polymer is %s\n", reactedPolymer)
	fmt.Printf("%d units\n", len(reactedPolymer))
}

func checkPair(pair string) string {
	newPair := ""
	first := float64(pair[0])
	second := float64(pair[1])
	if math.Abs(first-second) != 32 {
		newPair = pair
	}
	return newPair
}

func react(polymer string) string {
	var reactedPolymer string
	for i := 0; i < len(polymer)-1; i++ {
		substr := string(polymer[i : i+2])
		smolymer := checkPair(substr)
		if smolymer == substr {
			reactedPolymer += string(substr[0])
		} else {
			i++
		}
	}
	reactedPolymer += string(polymer[len(polymer)-1]) // there is probably a better way of doing this

	if reactedPolymer != polymer {
		// if something changed, do it again until it stops reacting
		reactedPolymer = react(reactedPolymer)
	}
	return reactedPolymer
}
