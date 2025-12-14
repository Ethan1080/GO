package main

import (
	"errors"
	"fmt"
)

// Une fonction simple qui renvoie soit un résultat, soit une erreur
func division(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("on ne peut pas diviser par zéro")
	}

	return a / b, nil
}

func main() {
	var a, b int

	fmt.Println("Hello, world!")
	fmt.Println("Choisisez un nombre")
	fmt.Scanln(&a)
	fmt.Println("Choisisez un autre nombre")
	fmt.Scanln(&b)

	// On essaye une division normale
	result, err := division(a, b)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println(a, " / ", b, " =", result)
	fmt.Println("Fin")
}
