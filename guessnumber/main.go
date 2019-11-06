package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	tries := 0
	score := 5
	number := randomizer()
	fmt.Println("Bienvenido al juego de adivinar número")
	game(number, tries, score)
}

func randomizer() int {
	rand.Seed(time.Now().UnixNano())
	max := 10
	return (rand.Intn(max) + 1)
}

func game(createdNumber int, tries int, score int) {
	inputnumber := 0
	if score <= 0 {
		fmt.Println("Perdiste, eres un estúpido")
		os.Exit(1)
	}
	if tries == 0 {
		fmt.Println("Listo, he pensado en un numero, dime tu suposición:")
	} else {
		fmt.Println("Intenta de nuevo:")
	}
	fmt.Scanf("%d", &inputnumber)
	if inputnumber < createdNumber {
		fmt.Println("Tu nùmero es menor al que pensé")
		tries++
		score = 5 - tries
		game(createdNumber, tries, score)
	}
	if inputnumber > createdNumber {
		fmt.Println("Tu nùmero es mayor al que pensé")
		tries++
		score = 5 - tries
		game(createdNumber, tries, score)
	}
	if inputnumber == createdNumber {
		fmt.Println("Felicidades has ganado, tu score es", score)
	}
}
