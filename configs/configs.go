package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	PublicHost  string
	Port        string
	// DBUser                  string
	// DBPassword              string
	// DBName                  string
	// DBAddress               string
	// DbServiceHost           string
	// PdfServiceHost          string
	// ValidatorServiceHost    string
	// CookiesAuthSecret       string
	// CookiesAuthAgeInSeconds int
	// CookiesAuthIsSecure     bool
	// CookiesAuthIsHttpOnly   bool
	// SessionSecret           string
	// GithubClientKey         string
	// GithubClientSecret      string
	// GithubClientCallback    string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Environment: getEnv("GO_ENV", "production"),
		PublicHost:  getEnv("PUBLIC_HOST", "http://localhost"),
		Port:        getEnv("PORT", "7331"),
		// DBUser:                  getEnv("DB_USER", "postgres"),
		// DBPassword:              getEnv("DB_PASSWORD", "postgres"),
		// DbServiceHost:           getEnv("DB_SERVICE_HOST", "localhost"),
		// PdfServiceHost:          getEnv("PDF_SERVICE_HOST", "localhost"),
		// ValidatorServiceHost:    getEnv("VALIDATOR_SERVICE_HOST", "localhost"),
		// DBName:                  getEnv("DB_NAME", "postgres"),
		// DBAddress:               fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", getEnv("DB_USER", "postgres"), getEnv("DB_PASSWORD", "postgres"), getEnv("DB_HOST", "127.0.0.1"), getEnvAsInt("DB_PORT", 5432), getEnv("DB_NAME", "postgres")),
		// CookiesAuthSecret:       getEnv("COOKIES_AUTH_SECRET", "some-very-secret-key"),
		// CookiesAuthAgeInSeconds: getEnvAsInt("COOKIES_AUTH_AGE_IN_SECONDS", twoDaysInSeconds),
		// CookiesAuthIsSecure:     getEnvAsBool("COOKIES_AUTH_IS_SECURE", false),
		// CookiesAuthIsHttpOnly:   getEnvAsBool("COOKIES_AUTH_IS_HTTP_ONLY", false),
		// SessionSecret:           getEnv("SESSION_SECRET", ""),
		// GithubClientKey:         getEnvOrError("GITHUB_KEY"),
		// GithubClientSecret:      getEnvOrError("GITHUB_SECRET"),
		// GithubClientCallback:    getEnvOrError("GITHUB_CALLBACK"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvOrError(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic(fmt.Sprintf("Environment variable %s is not set", key))
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}

		return b
	}

	return fallback
}
