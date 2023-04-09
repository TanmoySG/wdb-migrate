package config

import (
	"encoding/base64"
	"encoding/json"

	"github.com/TanmoySG/wunderDB/pkg/fs"
)

type Configurations struct {
	ConnectionConfigurations ConnectionConfigurations `json:"configurations"`
}

type ConnectionConfigurations struct {
	Retro    WdbRetroConfig `json:"retro"`
	Wunderdb WdbConfig      `json:"wunderdb"`
}

type WdbRetroConfig struct {
	BaseURL string               `json:"baseUrl"`
	Cluster EncodedConfiguration `json:"cluster"`
	Token   EncodedConfiguration `json:"token"`
}

type WdbConfig struct {
	BaseURL  string               `json:"baseUrl"`
	Username EncodedConfiguration `json:"username"`
	Password EncodedConfiguration `json:"password"`
}

type EncodedConfiguration string

func (ec EncodedConfiguration) Decode() string {
	decodedStringValue, err := base64.StdEncoding.DecodeString(string(ec))
	if err != nil {
		return ""
	}
	return string(decodedStringValue)
}

func LoadConfigurationsFromFile(configFilePath string) (*Configurations, error) {
	fileContentBytes, err := fs.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config Configurations
	err = json.Unmarshal(fileContentBytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
