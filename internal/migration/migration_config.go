package migration

import (
	"encoding/json"

	"github.com/TanmoySG/wunderDB/pkg/fs"
)

type MigrateDataConfig struct {
	Source SourceSink `json:"source"`
	Sink   SourceSink `json:"sink"`
}

type SourceSink struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
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

	return &config, nil
}
