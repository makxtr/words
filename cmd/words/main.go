package main

import (
	"fmt"
	"words/internal/app"
)

func main() {
	newApp, err := app.NewApp()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := newApp.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
