package migration

import (
	"encoding/json"
	"fmt"

	"github.com/TanmoySG/wunderDB/pkg/fs"
)

type MigrateDataConfig struct {
	Source SourceSink `json:"source"`
	Sink   SourceSink `json:"sink"`
}

type SourceSink struct {
	Database   *string `json:"database,omitempty"`
	Collection *string `json:"collection,omitempty"`
}

func LoadMigrationConfig(configFilePath string) (*MigrateDataConfig, error) {
	fileContentBytes, err := fs.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config MigrateDataConfig
	err = json.Unmarshal(fileContentBytes, &config)
	if err != nil {
		return nil, err
	}

	if config.Source.Collection == nil || config.Source.Database == nil {
		return nil, fmt.Errorf("configurations missing")
	}

	// setting default value for sink collection and db if not provided
	if config.Sink.Collection == nil || config.Sink.Database == nil {
		sinkCollection := fmt.Sprintf("migrated-%s", *config.Source.Collection)
		sinkDatabase := fmt.Sprintf("migrated-%s", *config.Source.Database)

		config.Sink.Database = &sinkDatabase
		config.Sink.Collection = &sinkCollection
	}

	return &config, nil
}
