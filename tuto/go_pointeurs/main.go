package main

import "fmt"

// &variable donne l'adresse de variable
// *variable va chercher a l'adresse variable

//avec les strucs ca le fait tt seul

func add1(x int) {
	//x = x + 1 //jm utilisé
}

func add1_pointeur(x *int) {
	*x = *x + 1
}

func main() {

	//1
	a := 10
	add1(a)
	fmt.Println("sans pointeur", a) // 10
	//

	//2
	b := 10
	add1_pointeur(&b)
	fmt.Println("avec pointeur", b) // 11
	fmt.Println(&b)                 //exemple: 0xc0000120e0
	//
}
