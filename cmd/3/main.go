package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getInput() string {
	return `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
}


func getInputFromFile() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)

}

func getLessCommonBit(line string) string {
    var number0 int
    var number1 int
    for i := 0; i < len(line); i ++ {
        number, _ := strconv.Atoi(string(line[i]))
        if number == 0 {
            number0++
        } else {
            number1++
        }
    }
    if number1 < number0 {
        return "1"
    } else {
        return "0"
    }
}

func getCommonBit(line string) string {
    var number0 int
    var number1 int
    for i := 0; i < len(line); i ++ {
        number, _ := strconv.Atoi(string(line[i]))
        if number == 0 {
            number0++
        } else {
            number1++
        }
    }
    if number1 > number0 {
        return "1"
    } else {
        return "0"
    }
}

func calculateDecimal(arrayIn []string) float64 {
    res := 0.00 
    for i := 0; i < len(arrayIn); i++ {
        if arrayIn[i] == "1" {
            res += math.Pow(2, float64(len(arrayIn)-i-1))
        }
    }
    return res 
}

func main() {
	var array []string

	data := strings.TrimSpace(getInputFromFile())
	parts := strings.Split(data, "\n")
	for _, part := range parts {
		array = append(array, part)
	}

	dimension := len(array[0])
    var result []string
    for i := 0; i < dimension; i++ {
	    for _, line := range array {
            result = append(result, string(line[i])) 
		}
	}
    var oppositeInput string
    for key, value := range result {

        if key % 12 == 0 && key > 0 {
            oppositeInput = oppositeInput + "\n"
        }

        oppositeInput = oppositeInput + value
    }

	lines := strings.Split(oppositeInput, "\n")
    var gammaBinary []string
    for _, line := range lines {
        gammaBinary = append(gammaBinary, getCommonBit(line))
    }

    var epsilonBinary []string
    for _, line := range lines {
        epsilonBinary = append(epsilonBinary, getLessCommonBit(line))
    }
    fmt.Println(gammaBinary)
    fmt.Println(epsilonBinary)
    fmt.Println(calculateDecimal(gammaBinary))
    fmt.Println(calculateDecimal(epsilonBinary))
}
