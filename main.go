package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ptrgags/mindless-stitchcraft/bracelets"
	"github.com/ptrgags/mindless-stitchcraft/patterns"
)

func zigzag(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: main.go knit-zigzag FABRIC_WIDTH MOTIF")
	}

	fabricWidth, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	motif := args[1]

	rows, err := patterns.GeneratePattern(motif, fabricWidth)

	if err != nil {
		return err
	}

	for _, row := range rows {
		fmt.Println(row)
	}

	return nil
}

func bracelet(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: main.go bracelet-repeat STRANDS MOTIF")
	}

	strandCount, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	motif, err := bracelets.ParseKnots(args[1])
	if err != nil {
		return err
	}

	rows, err := bracelets.GenerateUncoloredPattern(uint(strandCount), motif)
	if err != nil {
		return err
	}

	fmt.Println("Uncolored pattern:")
	for _, row := range rows {
		fmt.Println(row)
	}

	return nil
}

func main() {
	const usage = "usage: main.go {knit-zigzag,bracelet-repeat} ARGS"

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	var err error = nil
	switch os.Args[1] {
	case "knit-zigzag":
		err = zigzag(os.Args[2:])
	case "bracelet-repeat":
		err = bracelet(os.Args[2:])
	default:
		err = errors.New(usage)
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
