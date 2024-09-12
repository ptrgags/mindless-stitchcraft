package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ptrgags/mindless-knitting/patterns"
)

func main() {

	fmt.Println(os.Args)
	pattern := os.Args[2]
	fabricWidth, _ := strconv.Atoi(os.Args[3])

	rows, _ := patterns.GeneratePattern(pattern, fabricWidth)
	for _, row := range rows {
		fmt.Println(row)
	}

	fmt.Println()
	fmt.Println(pattern, fabricWidth)
}
