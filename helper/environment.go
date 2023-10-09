package helper

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Failed to load Environment Variable")
	}
}
