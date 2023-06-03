package main

import (
	"fmt"

	"sf_34.6.1/pkg/calc"
)

func main() {

	str, err := calc.OpenSave("./input.txt", "./output.txt")
	if err != nil {
		fmt.Println("Нет записей")
	}
	fmt.Println(string(str))

}
