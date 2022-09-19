// Package argparser package main
package argparser

import (
	"fmt"
	"os"
	"strconv"
)

type ArgParser struct {
	args []string
}

func NewArgParser(args []string) *ArgParser {
	return &ArgParser{args: args}
}

func (a ArgParser) Get() (int, int) {

	// Check for the correct number of arguments
	if len(a.args) != 2 {
		fmt.Println("null")
		os.Exit(1)
	}

	// Assign rows and bricks to variables. Ensure type
	rows, err := strconv.Atoi(a.args[0])
	if err != nil {
		fmt.Println("null")
		os.Exit(1)
	}

	bricks, err := strconv.Atoi(a.args[1])
	if err != nil {
		fmt.Println("null")
		os.Exit(1)
	}

	// Apply constraints
	if rows < 1 || bricks < 1 {
		fmt.Println("null")
		os.Exit(1)
	}

	if bricks > 10_000 {
		fmt.Println("Naah, too much...here's my resignation.")
		os.Exit(1)
	}

	return rows, bricks
}
