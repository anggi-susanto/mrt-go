package config

type Config struct {
	MongoConfig MongoConfig
}

type MongoConfig struct {
	Uri                  string
	Database             string
	WasteWaterCollection string
}
