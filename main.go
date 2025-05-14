package main

import (
	"fmt"
	"yourmodule/validator"
)

func main() {
	sgf := "(;GM[1]FF[4]SZ[19]AP[MyGoApp:1.0];B[dd];W[dc];B[cd];W[cc])"
	err := validator.ValidateSGF(sgf)
	if err != nil {
		fmt.Println("Ошибка валидации:", err)
	} else {
		fmt.Println("Все ходы допустимы")
	}
}
