package main

import (
	"os"
	art "ascii/functions"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	// args := os.Args[1:]
	outputFlag := flag.String("output", "output.txt", "Output file name")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 || len(args) > 3 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if args[0] == "" {
		return
	}
	if args[0] == "\\n" {
		fmt.Println()
		return
	}
	// Validation and processing
	var banner, input, output string
	output = *outputFlag

	if len(args) == 1 {
		input = args[0]
		banner = "standard"
     	result := art.ProcessFile(banner, input)
		fmt.Println(result)

	} else if len(args) == 2 {
		input = args[0]
		if strings.HasSuffix(args[1], ".txt") {
			banner = strings.TrimSuffix(args[1], ".txt")
		} else {
			banner = strings.ToLower(args[1])
		}
		result := art.ProcessFile(banner, input)
		fmt.Println(result)
	} else if len(args) == 3 {
		if !strings.HasPrefix(args[0], "--output=") {
			log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		}
		output = strings.TrimPrefix(args[0], "--output=")
		if !strings.HasSuffix(output, ".txt") {
			log.Fatal("wrong file extension: file must have a '.txt' extension")
		}
		input = args[1]
		banner = strings.ToLower(args[2])
	} else {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if len(output) < 4 {
		log.Fatal("Error: The .txt file size should be greater than 4")
	}
	if output == "shadow.txt" || output == "standard.txt" || output == "thinkertoy.txt" {
		log.Fatal("The name of the text file is the same as the banner file name. Please use a different file name.")
	}
	if strings.Count(output, ".") > 1 {
		log.Fatal("Usage: filename.txt")
	}

	// Process the file
	result := art.ProcessFile(banner, input)
	os.WriteFile(output,[]byte(result),0o644)
	// fmt.Println(result)
}
