package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type claim struct {
	claim_id   string
	x          int
	y          int
	width      int
	height     int
	overlapped bool
}

type square struct {
	num_claims int
	claim_ids  []string
}

func main() {
	fmt.Println("advent of code day 3")

	var claims_list []claim

	file, _ := os.Open("./day3.input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var claim_string = scanner.Text()
		s := strings.Split(claim_string, " ")
		var claim_id = s[0]
		var coords = s[2]
		var size = s[3]

		c := strings.Split(coords, ",")
		var x, _ = strconv.Atoi(c[0])
		var y, _ = strconv.Atoi(strings.Replace(c[1], ":", "", -1))

		z := strings.Split(size, "x")
		var width, _ = strconv.Atoi(z[0])
		var height, _ = strconv.Atoi(z[1])

		new_claim := claim{claim_id: claim_id, x: x, y: y, width: width, height: height, overlapped: false}
		claims_list = append(claims_list, new_claim)
	}

	var grid = [1000][1000]square{}

	for _, elem := range claims_list {
		for i := elem.x; i < elem.x+elem.width; i++ {
			for j := elem.y; j < elem.y+elem.height; j++ {
				grid[i][j].num_claims += 1
				grid[i][j].claim_ids = append(grid[i][j].claim_ids, elem.claim_id)
			}
		}
	}

	var multi_claim_count = 0
	for i, row := range grid {
		for j, _ := range row {
			if grid[i][j].num_claims > 1 {
				multi_claim_count += 1
				for _, claim_id := range grid[i][j].claim_ids {
					mark_overlapped(claims_list, claim_id)
				}
			}
		}
	}

	fmt.Printf("%d grid squares have multiple claims on them\n", multi_claim_count)

	for _, elem := range claims_list {
		if elem.overlapped == false {
			fmt.Printf("Claim %s has no overlapping claims\n", elem.claim_id)
		}
	}
}

func mark_overlapped(claims_list []claim, claim_id string) {
	for i := 0; i < len(claims_list); i++ {
		if claims_list[i].claim_id == claim_id {
			claims_list[i].overlapped = true
		}
	}
}
