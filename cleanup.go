package main

import (
	"log"
	"time"

	"github.com/gar-r/ppnext/config"
)

func scheduleBackgroundCleanup() {

	ch := time.NewTicker(config.CleanupFrequency)
	go func() {
		for range ch.C {
			err := config.Repository.Cleanup(config.MaximumRoomAge)
			if err != nil {
				log.Println(err)
			}
		}
	}()

}
