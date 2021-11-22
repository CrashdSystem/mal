package main

import (
	"bufio"
	"fmt"
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
		s, _ := reader.ReadString('\n')
		fmt.Println(rep(s))
	}
}
