package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

	runes := []rune(s)
	slices.Reverse(runes)

	fmt.Printf("\n%s -> %s \n\n", s, string(runes))

	return s == string(runes)
}

func main() {
	fmt.Println("------ Palíndromo ------")
	fmt.Println("> Informe sua frase:")

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		phrase := scanner.Text()

		if isPalindrome(phrase) {
			fmt.Println("É um palíndromo!! o/")
		} else {
			fmt.Println("Não é um palíndromo!!")
		}
	}

	fmt.Println("------------------------")
}
