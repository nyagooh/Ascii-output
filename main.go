package main

import (
	"fmt"

	// "log"
	"os"
	"strings"

	art "ascii/functions"
)

func main() {
	args := os.Args[1:]
	if len(args) > 2 || len(args) == 0 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	if args[0] == "" {
		return
	}
	if args[0] == "\\n" {
		fmt.Println()
		return
	}
	var banner string
	if len(args) < 2 {
		banner = "standard"
	} else if len(args) == 2 && strings.HasSuffix(args[1], ".txt") {
		new := strings.Trim(args[1], ".txt")
		banner = new
	} else if !strings.HasSuffix(args[1], ".txt") && len(args) == 2 {
		banner = args[1]
	}
	filename := fmt.Sprintf("%s%s", banner, ".txt")
	input := args[0]
	printable := art.NonPrintable(input)
	lines := strings.Split(printable, "\\n")
	for _, line := range lines {
		if line != "" {
			result := art.ProcessLine(line)
			result2 := art.ProcessString(result, filename)
			art.PrintStrings(result2)
		} else {
			fmt.Println()
		}
	}
}
