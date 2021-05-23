package utils

// Config : maps config variables to env
type Config struct {
	Environment string `envconfig:"ENVIRONMENT"`
	LogLevel    string `envconfig:"LOG_LEVEL"`
}
