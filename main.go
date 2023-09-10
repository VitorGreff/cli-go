package main

import (
	"cli/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world")
	app := app.Gerar()
	if err := app.Run(os.Args); err != nil {
		log.Fatal()
	}
}
