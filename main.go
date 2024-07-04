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
	if len(args) < 1 || len(args) == 3 {
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
	var input string
	var output string
	if len(args) < 2 && !strings.HasPrefix(args[0], "--output=") {
		banner = "standard"
		output = "output.txt"
		input = args[0]

	} 
	if len(args) == 2 && strings.HasSuffix(args[1], ".txt") && !strings.HasPrefix(args[0], "--output=") {
		new := strings.Trim(args[1], ".txt")
		banner = strings.ToLower(new)
		input = args[0]
		output = "output.txt"
	}
	if !strings.HasSuffix(args[1], ".txt") && len(args) == 2 && !strings.HasPrefix(args[0], "--output=") {
		banner = strings.ToLower(args[1])
		input = args[1]
	}

	if len(args) == 3 && strings.HasPrefix(args[0], "--output=") {
		output = args[0][9:]
		if !strings.HasSuffix(output, ".txt") {
			fmt.Println("wrong file Extension: file must have a '.txt' extension")
		}
		input = args[1]
		banner = args[2]
	}
	if len(output) <= 4 {
		fmt.Println("The .txt file should be greater that four")
		return
	}
	if output == "shadow" || output == "standard" || output == "thinkertoy" {
		fmt.Println("Name of the txt file is the same as the banner file name : include another file name")
		return
	}
	if strings.Count(output, ".") > 1 {
		fmt.Println("Usage: filename.txt")
		return
	}
	filename := fmt.Sprintf("%s%s", banner, ".txt")
	printable := art.NonPrintable(input)
	lines := strings.Split(printable, "\\n")
	for _, line := range lines {
		if line != "" {
			result := art.ProcessLine(line)
			result2 := art.ProcessString(result, filename)
			str := art.PrintStrings(result2)
			err := os.WriteFile(output, []byte(str), 0o644)
			if err != nil {
				fmt.Println("error writing", err)
			}
		} else {
			err := os.WriteFile(output, []byte("\n"), 0o644)
			if err != nil {
				fmt.Println("error writing", err)
			}
			// fmt.Println()
		}
	}
}
