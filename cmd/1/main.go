package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	return `199
200
208
210
200
207
240
269
260
263`
}

func getInputFromFile() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)

}

func main() {
	previous := 0
	result := 0

	res := getInputFromFile()
	lines := strings.Split(res, "\n")
	for key, line := range lines {
		number, _ := strconv.Atoi(line)

		if key > 0 {
			if number > previous {
				result++
			}
			previous = number
		}

	}
	fmt.Println(result)
}
