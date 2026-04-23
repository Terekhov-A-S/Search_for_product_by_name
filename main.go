package main

import (
	"fmt"
	"strings"
)

func main() {
	var productName string
	fmt.Print("Введите название товара: ")
	fmt.Scanln(&productName)
	nameToLower := strings.ToLower(productName)

	switch {
	case strings.Contains(strings.ToLower("Клавиатура JZ9"), nameToLower):
		fmt.Println("Клавиатура JZ9: 19200")
	case strings.Contains(strings.ToLower("Наушники N45"), nameToLower):
		fmt.Println("Наушники N45: 9600")
	case strings.Contains(strings.ToLower("Смартфон S10"), nameToLower):
		fmt.Println("Смартфон S10: 55000")
	default:
		fmt.Printf("Товар \"%s\" не найден.\n", productName)
	}
}
