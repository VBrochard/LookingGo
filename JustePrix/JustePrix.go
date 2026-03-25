package justeprix

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
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

func Run() int {
	fmt.Println("Bienvenue dans le Juste Prix")
	var devinette int
	var prix int
	essai := 0
	prix = rand.Intn(25000) + 5000
	for {
		fmt.Println("Guess :")
		essai += 1
		fmt.Scan(&devinette)
		if devinette == prix {
			fmt.Printf("Bravo, en %d essais", essai)
			return -1
		} else {
			if devinette-prix < 0 {
				clearScreen()
				fmt.Println("C'est plus")
			} else {
				clearScreen()
				fmt.Println("C'est moins")
			}
		}
	}
}
