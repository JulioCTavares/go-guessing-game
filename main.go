package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing game")

	fmt.Println("Um numero aleatorio sera sorteado. O numero é um inteiro de 0 a 100")

	number := rand.Int64N(101)

	scanner := bufio.NewScanner(os.Stdin)

	guessings := [10]int64{}

	for i := range guessings {
		fmt.Println("Qual seu chute? ")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)

		guessInt, err := strconv.ParseInt(guess, 10, 64)

		if err != nil {
			fmt.Println("Voce digitou algo invalido. Tente novamente.")
			return
		}

		switch {
		case guessInt < number:
			fmt.Println("voce errou. O numero sorteado é maior que", guessInt)
		case guessInt > number:
			fmt.Println("voce errou. O numero sorteado é menor que", guessInt)
		case guessInt == number:
			fmt.Printf(
				"Parabens! Voce acertou! O numero era: %d\n"+
					"Voce acertou em %d tentativas \n"+
					"Essas foram suas tentativas: %v\n",
				number, i+1, guessings[:i],
			)
			return
		}

		guessings[i] = guessInt
	}

	fmt.Printf(
		"Infelizmente voce não acertou o numero sorteado, que era %d\n"+
			"Voce teve 10 tentativas \n"+
			"Essas foram suas tentativas: %v\n",
		number, guessings,
	)
}
