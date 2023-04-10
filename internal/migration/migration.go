package migration

import (
	"fmt"

	"github.com/TanmoySG/wdb-migrate/pkg/wdb"
	"github.com/TanmoySG/wdb-migrate/pkg/wdb/retro"
)

type MigrationClient struct {
	Clients WdbClients
}

type WdbClients struct {
	WdbClient   wdb.WdbAdapter
	RetroClient retro.WdbRetroClient
}

func NewMigrationClient(wdbAdapter *wdb.WdbAdapter, wdbRetro *retro.WdbRetroClient) (*MigrationClient, error) {
	if wdbAdapter == nil || wdbRetro == nil {
		return nil, fmt.Errorf("wdb clients nil")
	}
	return &MigrationClient{
		Clients: WdbClients{
			WdbClient:   *wdbAdapter,
			RetroClient: *wdbRetro,
		},
	}, nil
}
