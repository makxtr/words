package main

import (
	"fmt"
	"words/internal/app"
)

func main() {
	newApp := app.NewApp()
	if err := newApp.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
