package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("advent of code day 1")

	var freq int = 0

	m := make(map[int]int)
	i := 0

	for {
		file, _ := os.Open("./day1.input")
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			var change = scanner.Text()
			change_num, e := strconv.Atoi(change)
			if e != nil {
				log.Fatal(e)
			} else {
				freq += change_num

				if _, ok := m[freq]; ok {
					fmt.Printf("seeing freq %d AGAIN\n", freq)
					m[freq] += 1
					os.Exit(3)
				} else {
					m[freq] = 1
				}

			}
		}
		i += 1
		file.Close()
	}
}
