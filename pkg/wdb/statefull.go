package wdb

type wdbStateAdapter struct {
	client WdbAdapter
	State  CurrentState
}

type CurrentState struct {
	Database   string
	Collection *string
}

func (wdbA *WdbAdapter) Use(database string, collections ...string) wdbStateAdapter {
	var collectionName *string
	if len(collections) > 0 {
		collectionName = &collections[0]
	}

	return wdbStateAdapter{
		client: *wdbA,
		State: CurrentState{
			Database:   database,
			Collection: collectionName,
		},
	}
}
