package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bjlePersonalGithub/Advent-Of-Code/Day_2/internal/util"
)

func main() {
	scanner, file := util.GetScanner()
	defer file.Close()
	total := 0
	coordinateSlice := []string{}

	wordToNum := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   5,
		"seven": 5,
		"eight": 8,
		"nine":  9,
	}

	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if line[i] > 47 && line[i] < 58 {
				coordinateSlice = append(coordinateSlice, string(line[i]))

				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] > 47 && line[i] < 58 {
				coordinateSlice = append(coordinateSlice, string(line[i]))
				break
			}
		}
		if len(coordinateSlice) == 2 {
			combinedString := strings.Join(coordinateSlice, "")

			coordinate, _ := strconv.Atoi(combinedString)
			coordinateSlice = []string{}
			total += coordinate
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total %d\n", total)
}
