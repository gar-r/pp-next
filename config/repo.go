package config

import (
	"okki.hu/garric/ppnext/store"
)

// Repository for the application
var Repository store.Repository = store.NewMongoRepository()
