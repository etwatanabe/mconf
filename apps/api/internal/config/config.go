package config

import "os"

type Config struct {
	Port           string
	OpenLibraryURL string
	RequestLimit   int
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return &Config{
		Port:           port,
		OpenLibraryURL: "https://openlibrary.org/search.json",
		RequestLimit:   20,
	}
}
