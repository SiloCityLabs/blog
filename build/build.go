package main

import (
	"log"
	"os"
)

func main() {
	loadSettings()

	// Scan for changes from events.json file

	if settings.BuildFontello {
		if err := buildFontello(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	// if settings.BuildBootstrap {
	// 	if err := buildBootstrap(); err != nil {
	// 		log.Println(err)
	// 		os.Exit(1)
	// 	}
	// }

	if settings.BuildSite {
		if err := buildSite(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	// if settings.SendNewsletter {
	// 	sendNewsletter()
	// }
}
