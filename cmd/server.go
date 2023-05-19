package main

import (
	"log"
	"main/src/cmd"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
}
