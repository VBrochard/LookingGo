package main

import (
	justeprix "LookingGo/JustePrix"
	"LookingGo/pendu"
	"fmt"
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

func main() {
	for {
		clearScreen()
		fmt.Println("Menu de LookingGo")
		fmt.Println("1. Pendu")
		fmt.Println("2. Le Juste Prix")
		fmt.Println("0. Quiter")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 1:
			if pendu.Run() == -1 {
				continue
			}
		case 2:
			justeprix.Run()
		case 0:
			return
		}
	}
}
