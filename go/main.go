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

    parser := interpreter.Parser{interpreter.Tokenizer{string, 0, nil}, &interpreter.Token{}}
		interpreter := interpreter.Interpreter{}
    interpreter.AddParser(parser)

		result := interpreter.Interpret()
		fmt.Println("Result ", result)

	}

}
