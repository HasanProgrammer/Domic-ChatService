package InfrastructureConcrete

import (
	"github.com/json-iterator/go"
	"os"
	"path/filepath"
)

type Config struct {
	ConnectionStrings struct {
		IRabbitMQ  string `json:"I-RabbitMQ"`
		IRedis     string `json:"I-Redis"`
		PostgreSql string `json:"PostgreSql"`
		SqlServer  string `json:"SqlServer"`
	} `json:"ConnectionStrings"`
}

type Configuration struct{}

func (configuration *Configuration) GetConnectionString(key string) (string, error) {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	dir, err := os.Getwd()

	configFilePath := filepath.Join(dir, "src", "Presentation", "Domic.WebAPI", "Config.json")

	data, err := os.ReadFile(configFilePath)

	if err != nil {

	}

	var config *Config

	json.Unmarshal(data, &config)

	switch key {
	case "SqlServer":
		return config.ConnectionStrings.SqlServer, nil
	case "PostgreSql":
		return config.ConnectionStrings.PostgreSql, nil
	case "I-RabbitMQ":
		return config.ConnectionStrings.IRabbitMQ, nil
	case "I-Redis":
		return config.ConnectionStrings.IRedis, nil
	default:
		return "", nil
	}

}

func NewConfiguration() *Configuration {
	return &Configuration{}
}
