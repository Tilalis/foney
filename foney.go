package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"foney/interpreter"
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
		// (1$ / 3) + (5$ * 2)  -- fails because of trailing space
		// febf=300$\nfebf = febf - 200$ -- some fail in calculations
		scanner = bufio.NewScanner(strings.NewReader("a = 5\na = a*a + a"))
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	fmt.Print(prompt)
	for scanner.Scan() {
		input = scanner.Text()

		t := time.Now()
		result, err := interpreter.InterpretString(input)
		fmt.Printf("    %s\n", time.Since(t))

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		t = time.Now()
		resultVM, err := interpreter.InterpretStringVM(input)
		fmt.Printf("VM: %s\n", time.Since(t))

		if err != nil {
			fmt.Printf("%v\n", err)
		}

		if result != nil {
			fmt.Printf("    %v\n", result)
		}

		if resultVM != nil {
			fmt.Printf("VM: %v\n", resultVM)
		}

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
				lexer, err := interpreter.NewLexer(strings.NewReader(input))

				if err != nil {
					goto End
				}

				parser, err := interpreter.NewParser(lexer)

				if err != nil && !errors.Is(err, interpreter.ErrUnsupportedOperation) {
					goto End
				}

				node, err := parser.Parse()

				if err != nil {
					goto End
				}

				bytecode := &interpreter.Instruction{
					Instruction: interpreter.NOP,
				}

				_, err = node.Compile(bytecode)

				if err != nil {
					fmt.Printf("%v\n", err)
				}

				loaded := bytecode.Load()

				if err != nil {
					fmt.Printf("%v\n", err)
				} else {
					for _, bc := range loaded {
						if bc.Argument != nil {
							fmt.Printf("%v %v\n", bc.Instruction, bc.Argument)
						} else {
							fmt.Println(bc.Instruction)
						}
					}
				}
			}
		}

	End:
		fmt.Print(prompt)
	}
}
