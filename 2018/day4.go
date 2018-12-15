package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("advent of code day 4")

	file, _ := os.Open("./day4.input")
	var shift_lines = []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		shift_lines = append(shift_lines, scanner.Text())
	}

	sort.Strings(shift_lines)

	guards := make(map[string]*[60]int)
	var guard_id string
	var start_sleep, end_sleep int

	for _, elem := range shift_lines {
		var date_string = strings.Split(elem, " ")[1]
		date_string = strings.Replace(date_string, "]", "", -1)
		var event = elem[19:]

		// these are in order so one we see "Guard X starts shift", it's all about them until we see a new guard
		if event[:5] == "Guard" {
			_, err := fmt.Sscanf(event, "Guard %s begins shift", &guard_id)
			if err != nil {
				log.Fatal(err)
			}
		} else if event == "falls asleep" {
			start_sleep, _ = strconv.Atoi(strings.Split(date_string, ":")[1])
		} else if event == "wakes up" {
			end_sleep, _ = strconv.Atoi(strings.Split(date_string, ":")[1])
			if guards[guard_id] == nil {
				guards[guard_id] = &[60]int{}
			}
			for min := start_sleep; min < end_sleep; min++ {
				guards[guard_id][min] += 1
			}
		}
	}

	max_minutes_asleep := 0
	var sleepiest_guard string
	for guard_id, guard := range guards {
		mins_asleep := 0
		for _, min := range guard {
			mins_asleep += min
		}
		if mins_asleep > max_minutes_asleep {
			sleepiest_guard = guard_id
			max_minutes_asleep = mins_asleep
		}
	}
	fmt.Printf("The sleepiest guard was %s for %d minutes\n", sleepiest_guard, max_minutes_asleep)

	var sleepiest_min int  // index / minute number
	var sleepiest_slep int // number of minutes
	for i, min := range guards[sleepiest_guard] {
		if min > sleepiest_slep {
			sleepiest_min = i
			sleepiest_slep = min
		}
	}
	fmt.Printf("The sleepiest minute for guard %s was %d with %d sleps\n", sleepiest_guard, sleepiest_min, sleepiest_slep)

	sleepiest_id_int, _ := strconv.Atoi(strings.Replace(sleepiest_guard, "#", "", -1))
	fmt.Println(sleepiest_id_int * sleepiest_min)

	var sleepiestest_min int  // index
	var sleepiestest_slep int // minutes
	var sleepiestest_guard string
	for guard_id, guard := range guards {
		for i, min := range guard {
			if min > sleepiestest_slep {
				sleepiestest_min = i
				sleepiestest_guard = guard_id
				sleepiestest_slep = min
			}
		}
	}
	fmt.Printf("Sleepiest guard ever is guard %s at minute %d\n", sleepiestest_guard, sleepiestest_min)
	sleepiestest_id_int, _ := strconv.Atoi(strings.Replace(sleepiestest_guard, "#", "", -1))
	fmt.Println(sleepiestest_id_int * sleepiestest_min)
	// do not @ me about these variable names. I know they are bad and I should feel bad.
}
