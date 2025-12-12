package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var entry int
	try := 0
	number := rand.Intn(100)
	fmt.Println("Devine un nombre entre 0 et 99")
	for {
		fmt.Scanln(&entry)
		try++
		time.Sleep(1 * time.Second)
		if entry == number {
			break
		} else if entry > number {
			fmt.Println("Plus petit...")
		} else {
			fmt.Println("plus grand...")
		}
	}
	fmt.Println("Bravo ! le nombre était: ", number, "tu l'as trouvé en ", try, " essais!")
	time.Sleep(2 * time.Second)
}
