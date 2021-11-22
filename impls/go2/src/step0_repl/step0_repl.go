package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func READ(s string) string {
	return s
}

func EVAL(ast, env string) string {
	return ast
}

func PRINT(expr string) string {
	return expr
}

func rep(str string) string {
	return PRINT(EVAL(READ(str), ""))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("user> ")
		s, err := reader.ReadString('\n')

		if err != io.EOF {
			fmt.Println(rep(s))
		} else {
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}
