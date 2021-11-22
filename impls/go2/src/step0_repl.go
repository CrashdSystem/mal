package main

import "fmt"

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
	var s string
	for {
		fmt.Print("user> ")
		fmt.Scanln(&s)
		fmt.Println(rep(s))
	}
}
