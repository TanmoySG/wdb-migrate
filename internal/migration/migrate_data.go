package migration

import (
	"github.com/go-playground/log/v8"

	"github.com/TanmoySG/wdb-migrate/pkg/gen"
)

const wdbRetroIdKey = "_id"

var logFormat = "Moving Data from [%s/%s] to [%s/%s]"

func (mc MigrationClient) Migrate(source SourceSink, sink SourceSink) {
	log.Infof(logFormat, *source.Database, *source.Collection, *sink.Database, *sink.Collection)

	err := mc.Clients.WdbClient.CreateDatabase(*sink.Database)
	if err != nil {
		log.Fatalf("Sink Error: %s, DB not created", err)
	}

	log.Infof("Database Created in Sink: %s", *sink.Database)

	res, err := mc.Clients.RetroClient.GetData(*source.Database, *source.Collection)
	if err != nil {
		log.Fatalf("Source Error: %s", err)
	}

	// TODO: Check if Collection exists only thnen create it
	// mc.Clients.WdbClient.GetCollection(sink.Database, sink.Collection)

	generatedJsonSchema, err := gen.GenerateJsonSchema(res.Data)
	if err != nil {
		log.Fatalf("Schema Generation Error: %s", err)
	}

	err = mc.Clients.WdbClient.CreateCollection(*sink.Database, *sink.Collection, generatedJsonSchema)
	if err != nil {
		log.Fatalf("Sink Error: %s, Collection not created", err)
	}

	log.Infof("Collection Created in Sink: %s", *sink.Collection)

	// externalize the GenSchema and CreateCollection along with TODO to separate func

	for r, rb := range res.Data {
		cleanedData := rb
		delete(cleanedData.(map[string]interface{}), wdbRetroIdKey)

		md := mc.Clients.WdbClient.Use(*sink.Database, *sink.Collection)
		err := md.AddData(rb)
		if err != nil {
			log.Error(err)
			continue
		}

		log.Noticef("moved data with key: %s", r)
	}
}
