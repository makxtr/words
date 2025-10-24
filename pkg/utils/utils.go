package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	var cmd *exec.Cmd

	if "windows" == runtime.GOOS {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func ReadIntAnswer(min, max int) int {
	var answer int
	for {
		_, err := fmt.Scan(&answer)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if answer < min || answer > max {
			fmt.Printf("Please choose a number between %d and %d.\n", min, max)
			continue
		}

		return answer
	}
}
