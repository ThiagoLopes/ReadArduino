package config

import "os"

func GetEnvDefault(key, fallback string) string{
	if env := os.Getenv(key); env != ""{
		return env
	}
	return fallback
}
