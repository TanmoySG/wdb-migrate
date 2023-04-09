package wdb

type wdbStateAdapter struct {
	Client wdbAdapter
	State  CurrentState
}

type CurrentState struct {
	Database   string
	Collection *string
}

func (wdbA *wdbAdapter) Use(database string, collections ...string) wdbStateAdapter {
	var collectionName *string
	if len(collections) > 0 {
		collectionName = &collections[0]
	}

	return wdbStateAdapter{
		Client: *wdbA,
		State: CurrentState{
			Database:   database,
			Collection: collectionName,
		},
	}
}
