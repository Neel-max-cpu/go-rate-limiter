package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	env := flag.String("env", "local", "Environment")
	flag.Parse()

	envFile := ".env." + *env

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Error loading %s file", envFile)
	}

	log.Println("Loaded environment:", *env)
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
