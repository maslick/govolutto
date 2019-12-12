package src

import "os"

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getPort() string {
	var port = getEnv("PORT", "8080")
	return ":" + port
}
