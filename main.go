package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Simple basic BSL Repl")
	fmt.Println("Type 'exit' to quit")

	for {
		fmt.Print(">> ")

		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Exiting REPL.")
			break
		}

		fmt.Println(input)
	}
}
