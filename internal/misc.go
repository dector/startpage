package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Version   string
	Redirects map[string]string
}

type ExternalConfig struct {
	Redirects map[string]string `yaml:"redirects"`
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Version: "1.1.1",
	}

	ext := LoadExternalConfig()
	config.Redirects = ext.Redirects

	return config, nil
}

func getDefaultConfig() *Config {
	return &Config{
		Version:   "1.1.1",
		Redirects: getDefaultRedirects(),
	}
}

func LoadExternalConfig() ExternalConfig {
	config := ExternalConfig{}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return config
	}

	configPath := filepath.Join(homeDir, ".config", "startpage", "config.yml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Config file not found")
		return config
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config
	}

	var ext ExternalConfig
	err = yaml.Unmarshal(data, &ext)
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return config
	}

	return ext
}

func getDefaultRedirects() map[string]string {
	return map[string]string{
		"mail":   "https://gmail.com",
		"gmail":  "https://gmail.com",
		"chat":   "https://chatgpt.com",
		"chatc":  "https://claude.ai",
		"claude": "https://claude.ai",
		"yt":     "https://youtube.com",
	}
}
