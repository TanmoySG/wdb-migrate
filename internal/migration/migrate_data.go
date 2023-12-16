package migration

import (
	"log"

	"github.com/TanmoySG/wdb-migrate/pkg/gen"
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
		log.Fatalf("Source Error: %s", err)
	}

	// TODO: Check if Collection exists only thnen create it
	// mc.Clients.WdbClient.GetCollection(sink.Database, sink.Collection)

	generatedJsonSchema, err := gen.GenerateJsonSchema(res.Data)
	if err != nil {
		log.Fatalf("Schema Generation Error: %s", err)
	}

	err = mc.Clients.WdbClient.CreateCollection(sink.Database, sink.Collection, generatedJsonSchema)
	if err != nil {
		log.Fatalf("Sink Error: %s, Collection not created", err)
	}

	// externalize the GenSchema and CreateCollection along with TODO to separate func

	for r, rb := range res.Data {
		cleanedData := rb
		delete(cleanedData.(map[string]interface{}), wdbRetroIdKey)

		md := mc.Clients.WdbClient.Use(sink.Database, sink.Collection)
		err := md.AddData(rb)
		if err != nil {
			log.Println(err)
			continue
		}

		log.Printf("moved data with key: %s\n", r)
	}
}
