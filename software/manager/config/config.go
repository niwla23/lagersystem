package config

import "os"

var StoragePath string
var DBUri string
var TypesenseUri string
var OperatorBaseUrl string

func LoadConfigFromEnvironment() {
	StoragePath = os.Getenv("STORAGE_PATH")
	if StoragePath == "" {
		panic("STORAGE_PATH not set")
	}

	DBUri = os.Getenv("DB_URI")
	if DBUri == "" {
		panic("DB_URI not set")
	}

	TypesenseUri = os.Getenv("TYPESENSE_URI")
	if TypesenseUri == "" {
		panic("TYPESENSE_URI not set")
	}

	OperatorBaseUrl = os.Getenv("OPERATOR_BASE_URL")
	if TypesenseUri == "" {
		panic("OPERATOR_BASE_URL not set")
	}
}
