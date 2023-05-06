package main

import (
	util "ArenalSoundAlert/util"
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Initializing loop...")

		mediasApi := util.GetMediasFromApi()
		mediaDb := util.GetMediaDb()

		for _, mediaApi := range mediasApi {
			if mediaApi.Date.After(mediaDb.Date) {
				fmt.Println("New media found")
				util.SendMail(mediaApi)
				//util.LoadInDatabase(mediaApi)
			}
		}
		fmt.Println("Sleeping for 1min")
		time.Sleep(1 * time.Minute)
	}
}
