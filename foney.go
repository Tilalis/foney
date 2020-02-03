package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Tilalis/foney/compiler"
	"github.com/Tilalis/foney/interpreter"
	"github.com/Tilalis/foney/vm"
)

func main() {
	const prompt = "foney> "

	var (
		debug   *string = flag.String("debug", "", "debug")
		inDebug *bool   = flag.Bool("indebug", false, "indebug")
		input   string
	)

	flag.Parse()

	var scanner *bufio.Scanner

	if *inDebug {
		scanner = bufio.NewScanner(strings.NewReader("a = 1USD\na = a * 3"))
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	fmt.Print(prompt)
	for scanner.Scan() {
		input = scanner.Text()

		// INTERPRETER
		t := time.Now()
		result, err := interpreter.InterpretString(input)

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		if result != nil {
			fmt.Printf("    %v\n", result)
		}

		fmt.Printf("    %s\n", time.Since(t))

		// COMPILER + VM
		t = time.Now()

		code, err := compiler.CompileString(input)

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		result, err = vm.Execute(code)

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		if result != nil {
			fmt.Printf("VM: %v\n", result)
		}

		fmt.Printf("VM: %s\n", time.Since(t))

		// END

		if *inDebug || *debug != "" {
			if *inDebug || *debug == "lexer" || *debug == "all" {
				lexer, err := interpreter.NewLexer(strings.NewReader(input))

				if err == nil {
					for token, _ := lexer.Next(); token != nil; token, _ = lexer.Next() {
						fmt.Println(token)
					}
				}

				fmt.Println()
			}

			if *inDebug || *debug == "parser" || *debug == "all" {
				for _, codeItem := range code {
					fmt.Printf("%v %v\n", codeItem.Instruction, codeItem.Argument)
				}
			}
		}

		fmt.Print(prompt)
	}
}
