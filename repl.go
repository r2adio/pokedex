package main

import (
	"bufio"
	"fmt"
	"os"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(" > ")

		scanner.Scan()
		text := scanner.Text()

		fmt.Printf("echoing: %v\n", text)
	}
}
