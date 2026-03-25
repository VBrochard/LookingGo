package pendu

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"slices"
	"strings"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

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

func difficulty() int {
	clearScreen()
	fmt.Println("Bienvenue au jeu du pendu")
	fmt.Println("Choisissez une difficulté :")
	fmt.Println("1. Facile :")
	fmt.Println("2. Normal:")
	fmt.Println("3. Difficile :")
	fmt.Println("0. Retour menu :")
	var diff int
	fmt.Scan(&diff)
	switch diff {
	case 1:
		return 999
	case 2:
		return 15
	case 3:
		return 5
	case 0:
		return -1
	}
	return 0
}

func Run() int {
	dico := []string{"bonjour", "golang", "pendu", "clavier", "ornithorynque", "train", "processeur", "chaton", "zebre", "yoga"}
	mot := dico[rand.Intn(len(dico))]
	devinator := []string{}
	for range len(mot) {
		devinator = append(devinator, "_")
	}
	var devinette string
	var life int
	life = difficulty()
	if life == -1 {
		return -1
	}
	for {
		clearScreen()
		fmt.Printf("Il vous reste %d vies \n", life)
		fmt.Println(strings.Join(devinator, " "))
		fmt.Println("Donner une lettre :")
		fmt.Scan(&devinette)
		devinette = string(devinette[0])
		if !strings.Contains(mot, devinette) {
			life--
			if life == 0 {
				fmt.Println("Perdue")
				return 0
			}
		}
		devinator = research(devinette, mot, devinator)
		if !slices.Contains(devinator, "_") {
			fmt.Println("Le mot était bien " + strings.Join(devinator, ""))
			fmt.Println("Bravo c'est le gg")
			break
		}
	}
	return 0
}
