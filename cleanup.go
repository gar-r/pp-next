package main

import (
	"log"
	"time"

	"okki.hu/garric/ppnext/config"
	"okki.hu/garric/ppnext/consts"
)

func scheduleBackgroundCleanup() {

	ch := time.NewTicker(consts.CleanupFrequency)
	go func() {
		for range ch.C {
			log.Println(config.Repository.Cleanup(consts.MaximumRoomAge))
		}
	}()

}
