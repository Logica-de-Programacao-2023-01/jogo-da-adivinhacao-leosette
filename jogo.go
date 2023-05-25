package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	denovo := true
	jogadas := [][]int{}
	for denovo {
		numero := rand.Intn(100) + 1
		resp := ""
		tentativas := []int{}
		fmt.Println("Seja Bem-vindo ao jogo da adivinhação!")
		fmt.Println("Tente adivinhar o número entre 1 e 100.")

		for {
			fmt.Print("Digite um número: ")
			var palpite int
			_, err := fmt.Scan(&palpite)
			if err != nil {
				continue
			}

			tentativas = append(tentativas, palpite)

			if palpite == numero {
				fmt.Println("Parabéns, você acertou!")
				break
			} else if palpite > numero {
				fmt.Println("O número é menor do que", palpite)
			} else {
				fmt.Println("O número é maior do que", palpite)
			}
		}

		jogadas = append(jogadas, tentativas)

		fmt.Println("Você utilizou", len(tentativas), "tentativas.")

		fmt.Print("Deseja jogar novamente? (s/n): ")
		fmt.Scan(&resp)
		resp = strings.ToLower(resp)
		if resp != "s" {
			denovo = false
		}
	}

	fmt.Println("\nNúmero de tentativas em todas as jogadas:")
	for i, tentativas := range jogadas {
		fmt.Printf("Jogada %d: %d tentativas\n", i+1, len(tentativas))
	}
}
