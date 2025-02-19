package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string
	AppPort      string
	DBConnection string
	DBHost       string
	DBPort       int
	DBDatabase   string
	DBUserName   string
	DBPassword   string
	JWTSecret    string
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetString(key string, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)

	if err != nil {
		return fallback
	}

	return valAsInt
}

func GetBool(key string, fallback bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsBool, err := strconv.ParseBool(val)

	if err != nil {
		return fallback
	}

	return valAsBool
}

func GetConfig() Config {
	return Config{
		AppPort:      GetString("APP_PORT", ":8080"),
		AppName:      GetString("APP_NAME", "Daycare Preschool API"),
		DBConnection: GetString("DB_CONNECTION", "mysql"),
		DBHost:       GetString("DB_HOST", "localhost"),
		DBPort:       GetInt("DB_PORT", 3306),
		DBDatabase:   GetString("DB_DATABASE", "daycare"),
		DBUserName:   GetString("DB_USERNAME", "root"),
		DBPassword:   GetString("DB_PASSWORD", ""),
		JWTSecret:    GetString("JWT_SECRET", "secret"),
	}
}
