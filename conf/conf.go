package conf

import (
	"log"

	"github.com/ian-kent/gofigure"
)

type Config struct {
	StashPort  string `env:"STASH_PORT"                flag:"stash-port"                flagDesc:"Primary Stash Server Port"`
	Collection string `env:"MONGODB_COLLECTION"           flag:"mongodb-collection"           flagDesc:"MongoDB collection for data"`
	Database   string `env:"MONGODB_DATABASE"             flag:"mongodb-database"             flagDesc:"MongoDB database for data"`
	MongoDBURL string `env:"MONGODB_URL"                  flag:"mongodb-url"                  flagDesc:"MongoDB server URL"`
}

// DefaultConfig returns a pointer to a Config instance that has been populated
// with default values.
func DefaultConfig() Config {
	return Config{
		StashPort:  ":8080",
		Database:   "stash",
		Collection: "files",
		MongoDBURL: "mongodb://chs-mongo:27017",
	}
}

// Get returns a pointer to a config instance that has been populated with
// values provided by the environment or command-line flags, or with default
// values if none are provided.
func Get() *Config {
	cfg := DefaultConfig()
	err := gofigure.Gofigure(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
