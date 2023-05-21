package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	
	ignoreCase := flag.Bool("i", false, "Ignore case")
	invertMatch := flag.Bool("v", false, "Invert match")
	flag.Parse()


	pattern := flag.Arg(0)
	filePath := flag.Arg(1)


	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	
	scanner := bufio.NewScanner(file)

	
	caseSensitive := !*ignoreCase

	
	for scanner.Scan() {
		line := scanner.Text()
		match := false

		
		if caseSensitive {
			match = strings.Contains(line, pattern)
		} else {
			match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		}

		
		if *invertMatch {
			match = !match
		}

		
		if match {
			fmt.Println(line)
		}
	}

	
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
}
