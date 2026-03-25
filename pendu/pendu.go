package pendu

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

const (
	Reset = "\033[0m"
	Rouge = "\033[31m"
	Vert  = "\033[32m"
)

func loadDico() []string {
	nomFichier := "pendu/dico.txt"
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	cheminComplet := filepath.Join(exPath, nomFichier)
	file, err := os.Open(cheminComplet)
	if err != nil {
		return []string{"erreur", "fichier", "introuvable"}
	}
	defer file.Close()
	var mots []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mot := strings.TrimSpace(scanner.Text())
		if mot != "" {
			mots = append(mots, mot)
		}
	}
	return mots
}

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
				devinator[i] = Vert + string(lettre) + Reset
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
		return 12
	case 3:
		return 6
	case 0:
		return 0
	}
	return 0
}

func Run() int {
	dico := loadDico()
	mot := dico[rand.Intn(len(dico))]
	depot := []string{}
	devinator := []string{}
	for range len(mot) {
		devinator = append(devinator, Rouge+"_"+Reset)
	}
	var devinette string
	var life int
	life = difficulty()
	if life == -1 {
		return 0
	}
	for {
		clearScreen()
		fmt.Printf("Il vous reste %d vies \n", life)
		fmt.Println(strings.Join(devinator, " "))
		fmt.Println("Lettres déjà proposée :")
		fmt.Println(strings.Join(depot, " "))
		fmt.Println("Donner une lettre :")
		fmt.Scan(&devinette)
		devinette = string(devinette[0])
		if !strings.Contains(mot, devinette) {
			life--
			if !slices.Contains(depot, devinette) {
				depot = append(depot, devinette)
			}
			if life == 0 {
				fmt.Println("Perdue")
				return 0
			}
		}
		devinator = research(devinette, mot, devinator)
		if !slices.Contains(devinator, Rouge+"_"+Reset) {
			clearScreen()
			fmt.Println("Le mot était bien " + strings.Join(devinator, ""))
			fmt.Println("Bravo c'est le gg")
			fmt.Println("\n[Appuyez sur Entrée pour revenir au menu]")
			fmt.Scanln()
			break
		}
	}
	return 0
}
