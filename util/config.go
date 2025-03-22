package util

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	ServerPort string
	DBurl string
	JWTSigningKey string
}

// LoadConfig takes the path of a env file,
// and loads env variables into a config struct 
func LoadConfig(path string) (config Config, err error){
	err = godotenv.Load(path)
	if err != nil{
		return Config{}, err
	}

	ServerPort := os.Getenv("SERVER_PORT")
	DBurl := os.Getenv("DB_URL")
	JWTSigningKey := os.Getenv("JWT_SIGNING_KEY")





	return Config{
		ServerPort: ServerPort,
		DBurl: DBurl,
		JWTSigningKey: JWTSigningKey,
	}, nil
} 