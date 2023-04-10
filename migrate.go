package main

import (
	"log"
	"os"

	"github.com/TanmoySG/wdb-migrate/internal/config"
	"github.com/TanmoySG/wdb-migrate/internal/migration"
	"github.com/TanmoySG/wdb-migrate/pkg/wdb"
	"github.com/TanmoySG/wdb-migrate/pkg/wdb/retro"
)

const defaultConfigFilepath = "config.json"
const defaultSourceSinkConfigFilepath = "source-sink.json"

func main() {

	configFilePath := os.Getenv("CONFIG_FILEPATH")
	if configFilePath == "" {
		configFilePath = defaultConfigFilepath
	}

	c, err := config.LoadConfigurationsFromFile(configFilePath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	wdbrc := retro.NewClient(
		c.ConnectionConfigurations.Retro.BaseURL,
		c.ConnectionConfigurations.Retro.Cluster.Decode(),
		c.ConnectionConfigurations.Retro.Token.Decode(),
	)

	wdbac := wdb.NewClient(
		c.ConnectionConfigurations.Wunderdb.Username.Decode(),
		c.ConnectionConfigurations.Wunderdb.Password.Decode(),
		c.ConnectionConfigurations.Wunderdb.BaseURL,
	)

	sourceSinkConfigPath := defaultSourceSinkConfigFilepath
	if len(os.Args) > 1 {
		sourceSinkConfigPath = os.Args[1]
	}

	sourceSink, err := migration.LoadMigrationConfig(sourceSinkConfigPath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	mc, err := migration.NewMigrationClient(wdbac, &wdbrc)
	if err != nil {
		log.Fatalf(err.Error())
	}

	mc.MigrateData(sourceSink.Source, sourceSink.Sink)
}
