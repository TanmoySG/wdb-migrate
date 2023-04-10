package migration

import (
	"log"
)

const wdbRetroIdKey = "_id"

var logFormat = "Moving Data from [%s/%s] to [%s/%s]"

func (mc MigrationClient) MigrateData(source SourceSink, sink SourceSink) {
	log.Printf(
		logFormat,
		source.Database, source.Collection,
		sink.Database, sink.Collection,
	)

	res, err := mc.Clients.RetroClient.GetData(source.Database, source.Collection)
	if err != nil {
		log.Fatalf("Source Error: %s", err.Error())
	}

	for r, rb := range res.Data {
		cleanedData := rb
		delete(cleanedData.(map[string]interface{}), wdbRetroIdKey)

		md := mc.Clients.WdbClient.Use(sink.Database, sink.Collection)
		err := md.AddData(rb)
		if err != nil {
			log.Println(err.Error())
		}

		log.Printf("moved data with key: %s\n", r)
	}
}
