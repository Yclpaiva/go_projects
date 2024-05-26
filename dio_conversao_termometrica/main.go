package main

import "fmt"

func main() {
	var temperatura_em_c float64
	fmt.Print("Digite a temperatura em Celsius: ")
	fmt.Scanln(&temperatura_em_c)
	temperatura_em_k := temperatura_em_c + 273
	fmt.Println("A temperatura em Kelvin é: ", temperatura_em_k)
}
