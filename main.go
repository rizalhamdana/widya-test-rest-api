package main

import (
	"fmt"
	"os"

	configEnv "github.com/joho/godotenv"
)

func main() {

	err := configEnv.Load(".env")
	if err != nil {
		fmt.Println(".env is not loaded properly")
		os.Exit(2)
	}

	service := MakeHandler()
	service.HTTPServerMain()

}
