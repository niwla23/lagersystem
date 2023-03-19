package config

import "os"

var StoragePath string
var DBUri string

func LoadConfigFromEnvironment() {
	StoragePath = os.Getenv("STORAGE_PATH")
	if StoragePath == "" {
		panic("STORAGE_PATH not set")
	}

	DBUri = os.Getenv("DB_URI")
	if DBUri == "" {
		panic("DB_URI not set")
	}
}
