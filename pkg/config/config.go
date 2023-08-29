package config

import (
	"encoding/json"
	"fmt"
	"github-proxy/pkg/utils"
	"os"
)

type Source struct {
	Path string `json:"path"`
}

type AppConfig struct {
	Owner   string            `json:"owner"`
	Repo    string            `json:"repo"`
	Token   string            `json:"token"`
	Sources map[string]Source `json:"sources"`
}

var Config AppConfig

func LoadConfig() *AppConfig {
	file, err := os.ReadFile("./config.json")
	utils.Check(err)

	if err := json.Unmarshal(file, &Config); err != nil {
		fmt.Println("Failed to load config:")
		fmt.Println(string(file))
		panic(err)
	}

	return &Config
}

func ValidateConfig() {
	// utils.EnsureKeys(Config, []string{"token", "sources", "owner", "repo"})

	// sources := utils.EnsureMap(Config["sources"])

	// for _, val := range Config.Sources {
	//   endpoint := utils.EnsureMap(val)
	//   utils.EnsureKeys(endpoint, []string{"path"})
	// }
}

// func GetString(key string) string {
//   return Config[key].(string)
// }

// func GetSources() Sources {
//   return Config["sources"].(Sources)
// }
