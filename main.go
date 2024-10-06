package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/ptrgags/mindless-stitchcraft/bracelets"
	"github.com/ptrgags/mindless-stitchcraft/bracelets/repeat"
	"github.com/ptrgags/mindless-stitchcraft/knitting"
	"github.com/ptrgags/mindless-stitchcraft/knitting/sync"
	"github.com/ptrgags/mindless-stitchcraft/knitting/zigzag"
)

func knitZigzag(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: main.go knit-zigzag FABRIC_WIDTH MOTIF")
	}

	fabricWidth, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	motif, err := knitting.ParseMotif(args[1])
	if err != nil {
		return err
	}

	rows, err := zigzag.GenerateZigzagPattern(motif, fabricWidth)
	if err != nil {
		return err
	}

	for _, row := range rows {
		fmt.Println(row)
	}

	return nil
}

func knitSync(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: main.go knit-sync FABRIC_WIDTH MOTIF [MOTIF, ...]")
	}

	fabricWidth, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	motifStrs := args[1:]
	motifs := make([]knitting.Motif, len(motifStrs))
	for i, motifStr := range motifStrs {
		motif, err := knitting.ParseMotif(motifStr)
		if err != nil {
			return err
		}

		motifs[i] = motif
	}

	rows, err := sync.GeneratePattern(uint(fabricWidth), motifs)

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
		return errors.New("usage: main.go bracelet-repeat STRAND_LABELS MOTIF")
	}

	strandLabels := []rune(args[0])
	strandCount := len(strandLabels)

	motif, err := bracelets.ParseKnots(args[1])
	if err != nil {
		return err
	}

	rows, err := repeat.GenerateUncoloredPattern(uint(strandCount), motif)
	if err != nil {
		return err
	}

	fmt.Println("Uncolored pattern:")
	for _, row := range rows {
		fmt.Println(row)
	}

	coloredRows, err := repeat.GenerateColoredPattern(strandLabels, motif)
	if err != nil {
		return err
	}

	fmt.Println("Colored pattern:")
	for _, row := range coloredRows {
		fmt.Println(row)
	}

	return nil
}

func main() {
	const usage = "usage: main.go {knit-zigzag,knit-sync,bracelet-repeat} ARGS"

	if len(os.Args) < 2 {
		fmt.Println(usage)
		os.Exit(1)
	}

	var err error = nil
	switch os.Args[1] {
	case "knit-zigzag":
		err = knitZigzag(os.Args[2:])
	case "knit-sync":
		err = knitSync(os.Args[2:])
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
