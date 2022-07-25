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
	data := getInput()
	res := make([]int, 0)
	parsedData := strings.Split(data, "\n")
	for key, _ := range parsedData {
		if key < len(parsedData)-2 {
			first, _ := strconv.Atoi(parsedData[key])
			second, _ := strconv.Atoi(parsedData[key+1])
			third, _ := strconv.Atoi(parsedData[key+2])
			res = append(res, first+second+third)
		}

	}
    fmt.Println(res)
	previous := 0
	result := 0
	for key, number := range res {

		if key > 0 {
			if number > previous {
				result++
			}
			previous = number
		}

	}
	fmt.Println(result)
}
