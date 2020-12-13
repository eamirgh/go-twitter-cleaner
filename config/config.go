package config

import "os"

// Config is what's necessary for the app to run
type Config struct {
	AccessToken  string
	AccessSecret string
	APIKey       string
	APISecret    string
	ScreenName   string
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		AccessSecret: os.Getenv("ACCESS_SECRET"),
		APIKey:       os.Getenv("API_KEY"),
		APISecret:    os.Getenv("API_SECRET"),
		ScreenName:   os.Getenv("USERNAME"),
	}
}
