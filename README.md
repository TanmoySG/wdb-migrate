# wdb-migrate

Tool for wdb-retro `v0` to wdb `v1` data migration.

## Usage

To run the migration, generate the `config.json` and `source-sink.json` as mentioned [here](#connection-configuration). Then run the following command passing the source-sink json file path as a commandline argument.

```sh
go run migrate.go ./source-sink.json

# alternately run the make command
make run-migration
```

If source-sink.json is not provided, it picks up the default file at the project's root directory.

### Data Migration
<!-- 
Currently only data migration between pre-existing database is supported. So it is a prerequisite that the Database must exist in v1. Refer to [the issue `#1`](https://github.com/TanmoySG/wdb-migrate/issues/1) for more details on setting up the prerequisites. -->

Database and Collections would be created using generated schema from sampled data if it doesn't exist. If database/collection already exists in wdb v1 (sink) the process might fail.

### Connection Configuration

wdb-migrate uses a configuration JSON to get the required credentials to connect to v0 and v1 instances for the movement. The config json looks like

```json
{
    "configurations": {
        "retro": {
            "baseUrl": "",
            "cluster": "base64-encoded-value",
            "token": "base64-encoded-value"
        },
        "wunderdb": {
            "baseUrl": "",
            "username": "base64-encoded-value",
            "password": "base64-encoded-value"
        }
    }
}
```

The credentials - `cluster`, `token`, `username` and `password`, must be base64 encoded. To encode the credentials properly use the following shell command and use the same in the configuration json.

```sh
echo -n <value-to-encode> | base64
```

Refer to the [sample config json](./scripts/sample.config.json) for details.

### Migration Configuration

wdb-migrate uses a source-sink configuration json file to specify the source database & collection from where data needs to be migrated and sink database & collection to where data needs to be moved. The following should be the source-sink json file.

```json
{
    "source": {
        "database": "source-db",
        "collection": "source-collection"
    },
    "sink": {
        "database": "sink-db",
        "collection": "sink-collection"
    }
}
```

If sink configs are not provided, they're generated from using the source configs by prepending `migrated-` with the source entities.
