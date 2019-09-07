package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func buildFontello() {
	url := "http://fontello.com/"

	//Upload config file
	request, err := newfileUploadRequest(url, map[string]string{}, "config", "../themes/beautifulhugo/static/css/fontello-config.json")
	if err != nil {
		log.Println(err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bodyData, _ := ioutil.ReadAll(resp.Body)
	token := string(bodyData)

	//Lets download this file now
	out, _ := os.Create("/tmp/fontello-" + token + ".zip")
	defer out.Close()

	resp2, _ := http.Get(url + token + "/get")
	defer resp2.Body.Close()

	io.Copy(out, resp2.Body)

	//unzip Loop trhough and grab the files we need
	files, err := Unzip("/tmp/fontello-"+token+".zip", "/tmp/fontello-"+token)
	if err != nil {
		log.Println(err)
		return
	}

	//Delete the zip file, dont need it anymore
	os.Remove("/tmp/fontello-" + token + ".zip")

	for _, file := range files {

		if strings.Contains(file, "fontello-embedded.css") {
			os.Rename(file, "../themes/beautifulhugo/static/css/fontello-embedded.css")
		}

		if strings.Contains(file, "fontello.eot") {
			os.Rename(file, "../themes/beautifulhugo/static/font/fontello.eot")
		}

		if strings.Contains(file, "fontello.svg") {
			os.Rename(file, "../themes/beautifulhugo/static/font/fontello.svg")
		}
	}

	//Delete the folder, dont need it anymore
	os.RemoveAll("/tmp/fontello-" + token)

	//Update the site
	settings.BuildSite = true
}
