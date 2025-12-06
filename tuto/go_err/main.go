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
	// On essaye une division normale
	result, err := division(10, 2)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("10 / 2 =", result)

	// Maintenant, on provoque une erreur
	result, err = division(10, 0)
	if err != nil {
		fmt.Println("Erreur :", err) // <- affichera l’erreur
		return
	}
	fmt.Println("10 / 0 =", result)
}
