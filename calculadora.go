package main

import "fmt"

func main() {

	var (
		valor1   int
		sinal    string
		valor2   int
		continua string = "s"
	)

	for continua != "n" {
		println("---------------------------------------------------------")
		println("opcoes de calculadora: \nsoma='+' subtracao='-' divisao='/' multiplicacao='*'")
		println("---------------------------------------------------------")
		fmt.Scan(&valor1)
		fmt.Scan(&sinal)
		fmt.Scan(&valor2)
		if sinal == "+" {
			fmt.Println(valor1 + valor2)
		} else {
			if sinal == "-" {
				fmt.Println(valor1 - valor2)
			} else {
				if sinal == "/" {
					fmt.Println(valor1 / valor2)
				} else {
					if sinal == "*" {
						fmt.Println(valor1 * valor2)
					}
				}
			}
		}
		fmt.Println("deseja continuar[s/n]?")
		fmt.Scan(&continua)
	}
}