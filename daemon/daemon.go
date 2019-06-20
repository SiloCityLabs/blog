package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/go-playground/webhooks/github"
	"gopkg.in/yaml.v2"
)

func main() {

	daemon := flag.Bool("daemon", true, "bool")
	flag.Parse()

	loadSettings()

	if *daemon == false {
		updateSite()
		return
	}

	hook, _ := github.New(github.Options.Secret(settings.GithubSecret))

	http.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn't one of the ones asked to be parsed
				fmt.Printf("%+v\n", err.Error())
			}
		}
		switch payload.(type) {

		case github.PushPayload:
			fmt.Printf("Build executing...\n")
			go checkPush(payload.(github.PushPayload))
		}
	})
	http.ListenAndServe(":3000", nil)
}

// Check to see if the docs folder was updated, if so ignore becouse it was already built
func checkPush(pullRequest github.PushPayload) {
	for _, commit := range pullRequest.Commits {
		for _, file := range commit.Added {
			if strings.HasPrefix(file, "docs/") {
				return
			}
		}
		for _, file := range commit.Removed {
			if strings.HasPrefix(file, "docs/") {
				return
			}
		}
		for _, file := range commit.Modified {
			if strings.HasPrefix(file, "docs/") {
				return
			}
		}
	}

	//Docs hasnt been built. lets do it
	updateSite()
}

func updateSite() {

	// Git pull
	cmdPull := exec.Command("git", "pull")
	cmdPull.Dir = settings.RootPath
	cmdPull.Stdout = os.Stdout
	cmdPull.Stderr = os.Stderr
	if err := cmdPull.Run(); err != nil {
		fmt.Printf("cmdPull.Run() failed with %s\n", err)
	}

	// Image api minify
	//https://github.com/peterhellberg/tinypng

	// Max 1080p image

	// Delete folders to regenerate
	os.RemoveAll(settings.RootPath + "/docs/page/")
	os.RemoveAll(settings.RootPath + "/docs/categories/")
	os.RemoveAll(settings.RootPath + "/docs/post/")
	os.RemoveAll(settings.RootPath + "/docs/samples/")
	os.RemoveAll(settings.RootPath + "/docs/tags/")

	// Hugo build
	cmdBuild := exec.Command("hugo")
	cmdBuild.Dir = settings.RootPath
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	if err := cmdBuild.Run(); err != nil {
		fmt.Printf("cmdBuild.Run() failed with %s\n", err)
	}

	// Git Add
	cmdAdd := exec.Command("git", "add", ".")
	cmdAdd.Dir = settings.RootPath
	cmdAdd.Stdout = os.Stdout
	cmdAdd.Stderr = os.Stderr
	if err := cmdAdd.Run(); err != nil {
		fmt.Printf("cmdAdd.Run() failed with %s\n", err)
	}

	// Git Commit Message
	cmdCommit := exec.Command("git", "commit", "-m", "Webhook Rebuild")
	cmdCommit.Dir = settings.RootPath
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Stderr = os.Stderr
	if err := cmdCommit.Run(); err != nil {
		fmt.Printf("cmdCommit.Run() failed with %s\n", err)
	}

	// Git push
	cmdPush := exec.Command("git", "push", "origin", "master")
	cmdPush.Dir = settings.RootPath
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	if err := cmdPush.Run(); err != nil {
		fmt.Printf("cmdPush.Run() failed with %s\n", err)
	}

	// Newsletter trigger
}

type Settings struct {
	GithubSecret string
	TinyPNGKey   string
	Port         string
	RootPath     string
}

var settings Settings

func loadSettings() {
	file, err := os.Open("settings.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	yaml.Unmarshal(b, &settings)
}
