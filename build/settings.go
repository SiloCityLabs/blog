package main

import "flag"

//Settings used by this build program
type Settings struct {
	TinyPNG        string
	BuildSite      bool
	BuildFontello  bool
	BuildBootstrap bool
	SendNewsletter bool
}

var settings Settings

func loadSettings() {
	key := flag.String("tinypng", "", "tinypng api key")
	flag.Parse()

	settings.TinyPNG = *key

	//Temporary, lets always build the site
	settings.BuildSite = true
}
