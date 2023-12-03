package util

import (
	"bufio"
	"log"
	"os"
)

func GetScanner() (*bufio.Scanner, *os.File) {
	file, err := os.Open("data/coordinates.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner, file
}
