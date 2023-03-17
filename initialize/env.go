package initialize

import "github.com/joho/godotenv"

// LoadEnv loads the environment variables from the .env file

func LoadEnv(fileNames ...string) {
	err := godotenv.Load(fileNames[0])
	if err != nil {
		panic(err)
	}
}
