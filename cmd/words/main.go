package main

import (
	"fmt"
	"words/internal/app"
)

func main() {
	app := app.NewApp()
	if err := app.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
