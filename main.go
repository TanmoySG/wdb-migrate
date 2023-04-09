package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TanmoySG/wdb-migrate/internal/config"
	"github.com/TanmoySG/wdb-migrate/pkg/wdb/retro"
)

const defaultConfigFilepath = "config.json"

func main() {

	configFilePath := os.Getenv("CONFIG_FILEPATH")
	if configFilePath == "" {
		configFilePath = defaultConfigFilepath
	}

	c, err := config.LoadConfigurationsFromFile(configFilePath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	rc := retro.NewClient(
		c.ConnectionConfigurations.Retro.BaseURL,
		c.ConnectionConfigurations.Retro.Cluster.Decode(),
		c.ConnectionConfigurations.Retro.Token.Decode(),
	)

	res, err := rc.GetData("tsgOnWebData", "education")
	if err != nil {
		log.Fatalf(err.Error())
	}

	for r, rb := range res.Data {
		fmt.Printf("key: %s, value: %v\n\n", r, rb)
	}
}
