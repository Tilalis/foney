package main

import (
	"bufio"
	"fmt"
	"os"

	interpreter "foney/interpreter"
)

func main() {
	const prompt = "foney> "

	var (
		debug = false
		input string
	)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)

	for scanner.Scan() {
		input = scanner.Text()
		result, err := interpreter.InterpretString(input)

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		if result != nil {
			fmt.Printf("%v\n", result)
		}

		if debug {
			lexer, err := interpreter.NewLexer(input)
			if err == nil {
				for token, _ := lexer.Next(); token != nil; token, _ = lexer.Next() {
					fmt.Println(token)
				}
			}
		}

		fmt.Print(prompt)
	}
}
