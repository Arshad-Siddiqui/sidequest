package initialize

import "github.com/joho/godotenv"

// LoadEnv loads the environment variables from the .env file

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
