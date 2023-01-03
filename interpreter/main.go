package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ola",
		Usage: "Exibe a mensagem da sua primeira CLI",
		Action: func(*cli.Context) error {
			fmt.Println("Ola, minha primeira CLI!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
