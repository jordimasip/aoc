package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

func getInputFromFile() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	return string(b)

}

func main() {
	data := strings.TrimSpace(getInputFromFile())
	p := Point{
		x: 0,
		y: 0,
	}
	lines := strings.Split(data, "\n")
	res := 0
	for _, line := range lines {
		parts := strings.Split(line, " ")
		res, _ = strconv.Atoi(parts[1])
		if parts[0] == "forward" {
			p.x += res
		} else if parts[0] == "up" {
			p.y -= res
		} else if parts[0] == "down" {
			p.y += res
		}

	}
	fmt.Println(p.x * p.y)

}
