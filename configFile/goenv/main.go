package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	// godotenv package
	dotenv := goDotEnvVariable("STRONGEST_AVENGER")

	fmt.Printf("godotenv : %s = %s \n", "STRONGEST_AVENGER", dotenv)
}

// use os package to get the env variable which is already set
// func envVariable(key string) string {

//   // set env variable using os package
//   os.Setenv(key, "gopher")

//   // return the env variable using os package
//   return os.Getenv(key)
// }

// func main() {
//   // os package
//   value := envVariable("name")

//   fmt.Printf("os package: name = %s \n", value)
//   fmt.Printf("environment = %s \n", os.Getenv("APP_ENV"))
// }
