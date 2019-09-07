package main

func main() {
	loadSettings()

	// Scan for changes from events.json file

	if settings.BuildFontello {
		buildFontello()
	}

	if settings.BuildBootstrap {
		buildBootstrap()
	}

	if settings.BuildSite {
		buildSite()
	}

	if settings.SendNewsletter {
		sendNewsletter()
	}
}
