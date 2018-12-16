package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type gridSquare struct {
	closestPoint  int
	totalDistance int
}

func main() {
	fmt.Println("advent of code day 6")

	file, _ := os.Open("./day6.input")
	scanner := bufio.NewScanner(file)

	var maxX, maxY int
	var points []point

	for scanner.Scan() {
		line := scanner.Text()
		x, _ := strconv.Atoi(strings.Split(line, ", ")[0])
		y, _ := strconv.Atoi(strings.Split(line, ", ")[1])
		newPoint := point{x: x - 1, y: y - 1} // The - 1s are to account for the fact that the grid gets 0-indexed but the points are 1-indexed.
		// I know, that's wonky, but it makes the math work, just go with it.
		points = append(points, newPoint)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	var grid = make([][]gridSquare, maxY)
	for i := range grid {
		grid[i] = make([]gridSquare, maxX)
		for j := range grid[i] {
			grid[i][j].closestPoint = findClosestPoint(i, j, points, maxX+maxY)
			grid[i][j].totalDistance = findTotalDistance(i, j, points)
		}
	}

	// The "infinite" areas are ones that are on any edge of the grid.
	// that's everything in grid[0] and grid[maxX-1]
	var infinites = []int{}
	for i := range grid[0] {
		infinites = addInfinite(infinites, grid[0][i].closestPoint)
		infinites = addInfinite(infinites, grid[maxY-1][i].closestPoint)
	}
	for i := 0; i < maxY; i++ {
		infinites = addInfinite(infinites, grid[i][0].closestPoint)
		infinites = addInfinite(infinites, grid[i][maxX-1].closestPoint)
	}

	var biggestArea int
	for i := range points {
		if !isInList(infinites, i) {
			area := countArea(grid, i)
			if area > biggestArea {
				biggestArea = area
			}
		}
	}
	fmt.Printf("The biggest area is %d\n", biggestArea)

	var closestArea int
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j].totalDistance < 10000 {
				closestArea++
			}
		}
	}
	fmt.Printf("The closest area is %d squares big\n", closestArea)
}

func findClosestPoint(i int, j int, points []point, max int) int {
	var closestPoint int
	for index, point := range points {
		distance := getDistance(i, j, point)
		if distance < max {
			max = distance
			closestPoint = index
		} else if distance == max {
			// If it's a tie, -1 is how we'll indicate it belongs to no point
			closestPoint = -1
		}
	}

	return closestPoint
}

func findTotalDistance(i int, j int, points []point) int {
	var totalDistance int
	for _, point := range points {
		totalDistance += getDistance(i, j, point)
	}
	return totalDistance
}

func getDistance(i int, j int, p point) int {
	return int(math.Abs(float64(i-p.y)) + math.Abs(float64(j-p.x)))
}

func addInfinite(infinites []int, inf int) []int {
	var newInfinites []int
	var inList bool
	if inf == -1 {
		inList = true // not technically, but it makes this work
	} else {
		inList = isInList(infinites, inf)
	}
	if inList == false {
		newInfinites = append(infinites, inf)
	} else {
		newInfinites = infinites
	}

	return newInfinites
}

func isInList(list []int, item int) bool {
	inList := false
	for _, i := range list {
		if i == item {
			inList = true
			break
		}
	}
	return inList
}

func countArea(grid [][]gridSquare, item int) int {
	var area int

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j].closestPoint == item {
				area++
			}
		}
	}

	return area
}
