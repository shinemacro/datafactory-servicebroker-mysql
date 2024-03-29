package config

import (
	"encoding/json"

	"github.com/asiainfoLDP/broker_mysql/utils"
)

type Config struct {
	Port                       string `json:"port"`
	DataPath                   string `json:"data_path"`
	CatalogPath                string `json:"catalog_path"`
	ServiceInstancesFileName   string `json:"service_instances_file_name"`
	ServiceBindingsFileName    string `json:"service_bindings_file_name"`
	ServicdCredentialsFileName string `json:"service_credentials_file_name"`
}

var (
	currentConfiguration Config
)

func LoadConfig(path string) (*Config, error) {
	bytes, err := utils.ReadFile(path)
	if err != nil {
		return &currentConfiguration, err
	}

	err = json.Unmarshal(bytes, &currentConfiguration)
	if err != nil {
		return &currentConfiguration, err
	}
	return &currentConfiguration, nil
}

func GetConfig() *Config {
	return &currentConfiguration
}
