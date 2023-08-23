package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func Init() {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
		  log.Fatal("Error loading .env file")
		}
	})
}