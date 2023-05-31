package main

import (
	"fmt"

	"yandex-team.ru/bstask/app"
)

func main() {
	application, err := app.New()
	if err != nil {
		fmt.Printf("failed to initialize application: %v", err)
		return
	}
	fmt.Printf("application crashed: %v", application.Start())
}
