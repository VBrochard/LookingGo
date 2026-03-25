package main

import (
	"LookingGo/pendu"
	"fmt"
)

func main() {
	for {
		fmt.Println("1. Pendu")
		fmt.Println("0. Quiter")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			pendu.Run()
		case 0:
			return
		}
	}
}
