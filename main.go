package main

import (
	"fmt"

	// "log"
	"os"
	"strings"

	art "ascii/functions"
)

// var templates = map[string]string{
// 	"standard":   "standard.txt",
// 	"thinkertoy": "thinkertoy.txt",
// 	"shadow":     "shadow.txt",
// 	"best":       "best.txt",
// }

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
	if len(args) > 1 {
		banner = args[1]
	} else {
		banner = "standard"
	}
	filename := fmt.Sprintf("%s%s", banner, ".txt")
	input := args[0]
	printable := art.NonPrintable(input)
	lines := strings.Split(printable, "\n")
	result := art.ProcessLine(lines)
	result2 := art.ProcessString(result, filename)
	art.PrintStrings(result2)
}
