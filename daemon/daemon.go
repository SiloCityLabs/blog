package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/go-playground/webhooks/github"
)

var dir = "../"

func main() {

	daemon := flag.Bool("daemon", true, "bool")
	flag.Parse()

	if *daemon == false {
		updateSite()
		return
	}

	hook, _ := github.New(github.Options.Secret(getSecret()))

	http.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.PushEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
				fmt.Printf("%+v\n", err.Error())
			}
		}
		switch payload.(type) {

		case github.PushPayload:
			//pullRequest := payload.(github.PushPayload)
			//fmt.Printf("%+v\n", pullRequest)
			fmt.Printf("Build executing...\n")

			go updateSite()
		}
	})
	http.ListenAndServe(":3000", nil)
}

func getSecret() string {
	file, err := os.Open("secret.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	return string(b)
}

func updateSite() {

	// Git pull
	cmdPull := exec.Command("git", "pull")
	cmdPull.Dir = dir
	cmdPull.Stdout = os.Stdout
	cmdPull.Stderr = os.Stderr
	if err := cmdPull.Run(); err != nil {
		fmt.Printf("cmdPull.Run() failed with %s\n", err)
	}

	// Image api minify

	// Max 1080p image

	// Delete folders to regenerate
	os.RemoveAll("./docs/page/")
	os.RemoveAll("./docs/categories/")
	os.RemoveAll("./docs/post/")
	os.RemoveAll("./docs/samples/")
	os.RemoveAll("./docs/tags/")

	// Hugo build
	cmdBuild := exec.Command("hugo")
	cmdBuild.Dir = dir
	cmdBuild.Stdout = os.Stdout
	cmdBuild.Stderr = os.Stderr
	if err := cmdBuild.Run(); err != nil {
		fmt.Printf("cmdBuild.Run() failed with %s\n", err)
	}

	// Git Add
	cmdAdd := exec.Command("git", "add", ".")
	cmdAdd.Dir = dir
	cmdAdd.Stdout = os.Stdout
	cmdAdd.Stderr = os.Stderr
	if err := cmdAdd.Run(); err != nil {
		fmt.Printf("cmdAdd.Run() failed with %s\n", err)
	}

	// Git Commit Message
	cmdCommit := exec.Command("git", "commit", "-m", "Webhook Rebuild")
	cmdCommit.Dir = dir
	cmdCommit.Stdout = os.Stdout
	cmdCommit.Stderr = os.Stderr
	if err := cmdCommit.Run(); err != nil {
		fmt.Printf("cmdCommit.Run() failed with %s\n", err)
	}

	// Git push
	cmdPush := exec.Command("git", "push", "origin", "master")
	cmdPush.Dir = dir
	cmdPush.Stdout = os.Stdout
	cmdPush.Stderr = os.Stderr
	if err := cmdPush.Run(); err != nil {
		fmt.Printf("cmdPush.Run() failed with %s\n", err)
	}

	// Newsletter trigger
}
