package main

import (
	"fmt"
	"validate-go-moves/validator"
)

func main() {
	sgfContent := "(;GM[1]FF[4]SZ[19]AP[MyGoApp:1.0];B[dd];W[dc];B[cd];W[cc])"
	err := validator.ValidateSGF(sgfContent)
	if err != nil {
		fmt.Println("Ошибка валидации:", err)
	} else {
		fmt.Println("Все ходы допустимы")
	}
}
