package internal

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dector/kdly"
)

type Config struct {
	Version   string
	Redirects map[string]string
}

type ExternalConfig struct {
	Redirects map[string]string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Version: "1.2.1",
	}

	ext := LoadExternalConfig()
	config.Redirects = ext.Redirects

	return config, nil
}

func getDefaultConfig() *Config {
	return &Config{
		Version:   "1.1.1",
		Redirects: map[string]string{},
	}
}

func LoadExternalConfig() ExternalConfig {
	config := ExternalConfig{
		Redirects: make(map[string]string),
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return config
	}

	configPath := filepath.Join(homeDir, ".config", "startpage", "config.kdl")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Println("Config file not found")
		return config
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return config
	}

	doc, err := kdly.Parse(string(data))
	if err != nil {
		fmt.Println("Error parsing config file:", err)
		return config
	}

	// Find the "redirects" node
	for _, node := range doc.Nodes {
		if node.Name == "redirects" {
			// Parse children nodes as key-value pairs
			for _, child := range node.Children {
				if child.Name != "-" {
					continue
				}

				for _, arg := range child.Properties {
					config.Redirects[arg.Key] = arg.Value.Value
				}
			}
			break
		}
	}

	return config
}
