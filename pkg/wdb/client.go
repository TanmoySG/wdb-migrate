package wdb

import wdbgo "github.com/TanmoySG/wdb-go"

var projectID string = "wdb-migrate"

type wdbAdapter struct {
	wdbgo.Client
}

func NewClient(username string, password string, connectionURI string) (*wdbAdapter, error) {
	wdbG, err := wdbgo.NewClient(username, password, connectionURI, &projectID)
	if err != nil {
		return nil, err
	}

	return &wdbAdapter{wdbG}, nil
}
