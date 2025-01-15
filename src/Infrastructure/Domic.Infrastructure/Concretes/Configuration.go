package InfrastructureConcrete

import (
	"github.com/json-iterator/go"
	"os"
)

type Config struct {
	ConnectionStrings struct {
		IRabbitMQ  string `json:"I-RabbitMQ"`
		IRedis     string `json:"I-Redis"`
		PostgreSql string `json:"PostgreSql"`
	} `json:"ConnectionStrings"`
}

type Configuration struct{}

func (configuration *Configuration) GetConnectionString(key string) (string, error) {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	file, err := os.Open("Config.json")

	if err != nil {

	}

	defer file.Close()

	jsonDecoder := json.NewDecoder(file)

	var config *Config

	jsonDecoder.Decode(&config)

	switch key {
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
