package main

import "fmt"

func main()  {
	firstNumber := 0
	secondNumber := 0
	fmt.Print("Ingresa el primer digito: ")
	fmt.Scanf("%d", &firstNumber)
	fmt.Print("Ingresa el segundo digito: ")
	fmt.Scanf("%d", &secondNumber)
	fmt.Println("Resultado de la suma: ", sum(firstNumber,secondNumber))
	fmt.Println("Resultado de la resta: ", res(firstNumber,secondNumber))
}

func sum( a int, b int) int{
	return a + b
}

func res( a int, b int) int {
	return a - b
}