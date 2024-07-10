package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	DbTimezone string
	Fabric_path string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DbHost = os.Getenv("DB_HOST")
	fmt.Println("DbHost=================",DbHost)
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
	DbTimezone = os.Getenv("DB_TIMEZONE")
	Fabric_path = os.Getenv("FABRIC_PATH")
}

// err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	dbUser := os.Getenv("DB_USER")
// 	dbPassword := os.Getenv("DB_PASSWORD")
// 	dbName := os.Getenv("DB_NAME")
// 	dbTimezone := os.Getenv("DB_TIMEZONE")
