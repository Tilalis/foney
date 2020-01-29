package main

import (
	"bufio"
	"fmt"
	interpreter "github.com/Tilalis/foney.go/interpreter"
	money "github.com/Tilalis/foney.go/money"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("foney.go> ")
	for scanner.Scan() {
		input := scanner.Text()
		lexer, _ := interpreter.NewLexer(input)

		for token := lexer.Next(); token != nil; token = lexer.Next() {
			fmt.Println(token)
		}
		fmt.Print("foney.go> ")
	}

	input := "(5$ + Br15) / 2 - 5USD"
	lexer, _ := interpreter.NewLexer(input)

	for token := lexer.Next(); token != nil; token = lexer.Next() {
		fmt.Println(token)
	}

	usd, _ := money.GetCurrency("USD", "$")
	ten := money.NewMoney(12.3456, usd)

	fmt.Printf("%v", ten)

}
