package InfrastructureConcrete

import (
	"github.com/json-iterator/go"
	"os"
)

type Config struct {
	ConnectionStrings struct {
		PostgreSql string `json:"PostgreSql"`
	} `json:"ConnectionStrings"`
}

type Configuration struct{}

func (configuration *Configuration) GetPostgreSqlConnectionString(key string) (string, error) {

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
	default:
		return "", nil
	}

}

func NewConfiguration() *Configuration {
	return &Configuration{}
}
