package config

import "okki.hu/garric/ppnext/store"

var Repository store.Repository = store.NewMongoRepository() // Repository for the application
