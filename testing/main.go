package main

import (
	"fmt"
	"go-reloaded/fixer"
)

func main() {
	fmt.Println(fixer.BintoDecimal("1011"))
	fmt.Println(fixer.BintoDecimal("10"))

	fmt.Println(fixer.BintoDecimal("abc"))
}
