package wdb

import "fmt"

func (ws wdbStateAdapter) AddData(data any) error {
	if ws.State.Collection == nil {
		return fmt.Errorf("collection not set")
	}

	return ws.Client.AddData(data, ws.State.Database, *ws.State.Collection)
}