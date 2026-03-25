package pendu

import (
	"fmt"
	"slices"
	"strings"
)

func research(lettre string, mot string, devinator []string) []string {
	if strings.Contains(mot, lettre) {
		fmt.Println("Oui")
		for i := range len(mot) {
			if string(mot[i]) == lettre {
				devinator[i] = string(lettre)
			}
		}
		return devinator
	} else {
		fmt.Println("Non")
		return devinator
	}
}

func Run() {
	mot := "loupe"
	slice := []string{}
	for _, c := range mot {
		slice = append(slice, string(c))
	}
	devinator := []string{}
	for range len(mot) {
		devinator = append(devinator, "_")
	}
	var devinette string
	fmt.Println("Bienvenue au jeu du pendu")
	for {
		fmt.Println(strings.Join(devinator, " "))
		fmt.Println("Donner une lettre :")
		fmt.Scan(&devinette)
		devinette = string(devinette[0])
		devinator = research(devinette, mot, devinator)
		if !slices.Contains(devinator, "_") {
			fmt.Println("Le mot était bien " + strings.Join(devinator, ""))
			fmt.Println("Bravo c'est le gg")
			break
		}
	}
}
