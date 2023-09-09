package main

import (
	"bufio"
	"fmt"
	"igorwessel/rinha-compiler/interpreter"
	"log"
	"os"
)

func main() {

	for {

		fmt.Print("Digit you expression: ")
		reader := bufio.NewReader(os.Stdin)
		string, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal("noo")
		}

		interpreter := interpreter.Interpreter{interpreter.Tokenizer{string, 0, nil}, &interpreter.Token{}}

		result := interpreter.Expr()
		fmt.Println("Result ", result)

	}

}
