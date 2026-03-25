package justeprix

import (
	"fmt"
	"math/rand"
)

func Run() {
	fmt.Println("Bienvenue dans le Juste Prix")
	var devinette int
	var prix int
	essai := 0
	prix = rand.Intn(45000) + 5000
	for {
		fmt.Println("Guess :")
		essai += 1
		fmt.Scan(&devinette)
		if devinette == prix {
			fmt.Printf("Bravo, en %d essais", essai)
			break
		} else {
			if devinette-prix < 0 {
				fmt.Println("C'est plus")
			} else {
				fmt.Println("C'est moins")
			}
		}
	}
}
