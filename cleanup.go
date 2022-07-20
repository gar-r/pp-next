package main

import (
	"log"
	"time"

	"okki.hu/garric/ppnext/config"
)

func scheduleBackgroundCleanup() {

	ch := time.NewTicker(config.CleanupFrequency)
	go func() {
		for range ch.C {
			log.Println(config.Repository.Cleanup(config.MaximumRoomAge))
		}
	}()

}
