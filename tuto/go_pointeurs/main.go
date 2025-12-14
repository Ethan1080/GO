package main

import "fmt"

// &variable donne l'adresse de variable
// *variable va chercher a l'adresse variable

//avec les strucs ca le fait tt seul

func add1(x *int) {
	*x = *x + 1
}

func main() {
	a := 200
	var pointeur_exemple *int
	pointeur_exemple = &a
	*pointeur_exemple++
	fmt.Println()

	b := 10
	add1(&b)
	fmt.Println("avec pointeur", b)
	fmt.Println(&b)
}
