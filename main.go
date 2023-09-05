package main

import (
	"github.com/kamicodaxe/envzy-cli/cmd"
	"github.com/kamicodaxe/envzy-cli/internal/app"
)

func main() {
	_, err := app.InitializeDatabase()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
