package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const helloWorld = "Hola %s %s, bienvenido al fascinante mundo de Go.\n"

func main() {
	lastname := "<Modificar con mi apellido>"
	name := getName()
	var number = sum(50, 50)
	a, b, c := getVariables()
	f32, f64 := getFloat()
	stringUTF8 := getUnicode()
	fmt.Print("Ingresa tu apellido: ")
	fmt.Scanf("%s", &lastname)
	fmt.Printf(helloWorld, name, lastname)
	fmt.Println(number, lastname, a, b, c)
	fmt.Println(f32, f64)
	fmt.Println("Cadena UTF8", stringUTF8)
	fmt.Println(string("hola"[0]))
	fmt.Println("Cantidad de letras que tiene hola ", len("hola"), utf8.RuneCountInString("hola"))
	getArray()
	getSlice()
	ifTest()
	forTest()
	strings2()
}

func getName() string {
	var name string
	name = "Sin nombre"
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanf("%s", &name)
	return name
}

func getVariables() (int, int32, int64) {
	return 1, 2, 3
}

func sum(a int, b int) int {
	return a + b
}

func getFloat() (float32, float64) {
	return float32(0.1), float64(float32(0.1))
}

func getUnicode() string {
	return "こんにちは!"
}

func getArray() {
	var arr1 [2]string
	arr2 := [3]int{1, 2, 3}
	arr1[0] = "array"
	arr1[1] = "array2"
	fmt.Println(arr1)
	fmt.Println(arr2)
}

func getSlice() {
	var slice1 []string
	slice1 = append(slice1, "mi", "slice", "1")
	fmt.Println(slice1)
}

func ifTest() {
	var number = 0
	fmt.Print("Ingresa un número del 1 al 10: ")
	fmt.Scanf("%d", &number)
	if number%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}

	if number2 := 3; number2 == 3 {
		fmt.Println("Entró al condicional")
	}
}

func forTest() {
	for i := 0; i < 5; i++ {
		fmt.Println("FOR", i)
	}

	c := 100
	for c > 0 {
		c -= 10
		fmt.Println(c)
	}

	s := 10000
	for {
		s--
		if s == 0 {
			fmt.Println("Termina el for 'infinito'")
			break
		}
	}
}

func strings2() {
	var text = "Hello World, Hello Platzi, Hello Go"
	fmt.Println(strings.ToUpper(text))
	fmt.Println(strings.ToLower(text))
	fmt.Println(strings.Replace(text, "Hello", "Hola", -1))
}
