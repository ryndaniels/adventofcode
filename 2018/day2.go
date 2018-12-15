package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("advent of code day 2")

	var twices int = 0
	var thrices int = 0

	file, _ := os.Open("./day2.input")
	scanner := bufio.NewScanner(file)

	var all_ids []string

	for scanner.Scan() {
		var box_id = scanner.Text()
		all_ids = append(all_ids, box_id)
		if containsDouble(box_id) == true {
			twices += 1
		}
		if containsTriple(box_id) == true {
			thrices += 1
		}
	}

	var checksum int = twices * thrices
	fmt.Printf("checksum is %d\n", checksum)

	file.Close()

	file, _ = os.Open("./input.txt")
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var current_id = scanner.Text()
		for _, elem := range all_ids {
			if countDifferences(current_id, elem) == 1 {
				fmt.Println(commonLetters(current_id, elem))
				os.Exit(0)
			}
		}
	}
}

func containsDouble(id string) bool {
	var doubliness bool = false
	for _, elem := range id {
		if countInstances(id, string(elem)) == 2 {
			doubliness = true
		}
	}
	return doubliness
}

func commonLetters(str1 string, str2 string) string {
	b := strings.Builder{}
	for i, elem := range str1 {
		if string(elem) == string(str2[i]) {
			b.WriteString(string(elem))
		}
	}
	return b.String()
}

func countDifferences(str1 string, str2 string) int {
	var diffs int = 0
	for i, elem := range str1 {
		if string(elem) != string(str2[i]) {
			diffs += 1
		}
	}
	return diffs
}

func containsTriple(id string) bool {
	var tripliness bool = false
	for _, elem := range id {
		if countInstances(id, string(elem)) == 3 {
			tripliness = true
		}
	}
	return tripliness
}

func countInstances(id string, c string) int {
	var count int = 0
	for _, elem := range id {
		if string(elem) == c {
			count += 1
		}
	}
	return count
}
