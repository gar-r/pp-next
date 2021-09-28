package config

import (
	"okki.hu/garric/ppnext/consts"
	"okki.hu/garric/ppnext/store"
)

// Repository for the application
var Repository store.Repository = store.NewCache(
	store.NewFs(consts.RoomsPath),
)
