package wdb

import wdbgo "github.com/TanmoySG/wdb-go"

var projectID string = "wdb-migrate"

type WdbAdapter struct {
	wdbgo.Client
}

func NewClient(username string, password string, connectionURI string) *WdbAdapter {
	wdbG, err := wdbgo.NewClient(username, password, connectionURI, &projectID)
	if err != nil {
		return nil
	}

	return &WdbAdapter{wdbG}
}
