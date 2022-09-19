package main

import (
	"build-a-wall/cmd/argparser"
	"build-a-wall/cmd/printer"
	"fmt"
	"os"
)

func main() {
	// Read array of arguments
	args := argparser.NewArgParser(os.Args[1:])
	rows, bricks := args.Get()

	// Create printer
	printer := printer.NewPrinter(rows, bricks)

	// Get the message
	output := printer.GetOutput()

	// Print
	fmt.Print(output)
}
