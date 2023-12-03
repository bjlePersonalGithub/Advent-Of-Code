package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/bjlePersonalGithub/Advent-Of-Code/Day_1/internal/util"
)

func main() {
	scanner, file := util.GetScanner() // Capitalize the G in GetScanner
	defer file.Close()
	total := 0
	coordinateSlice := []string{}
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
