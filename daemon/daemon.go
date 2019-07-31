package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-playground/webhooks/github"
	"gopkg.in/yaml.v2"
)

func main() {

	rebuildFontello()
	os.Exit(1)

	daemon := flag.Bool("daemon", true, "bool")
	flag.Parse()

	loadSettings()

	if *daemon == false {
		updateSite()
		return
	}

	hook, _ := github.New(github.Options.Secret(settings.GithubSecret))

	http.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Server hit")

		payload, err := hook.Parse(r, github.PushEvent, github.PingEvent)
		if err != nil {
			fmt.Printf("%+v\n", err.Error())
			return
		}

		switch payload.(type) {

		case github.PushPayload:
			fmt.Printf("Build executing...\n")
			go checkPush(payload.(github.PushPayload))
		case github.PingPayload:
			fmt.Printf("Ping ... Pong\n")
		}
	})
	http.ListenAndServe(":"+settings.Port, nil)
}

// Check to see if the docs folder was updated, if so ignore becouse it was already built
func checkPush(pullRequest github.PushPayload) {
	//Lets only check first commit, just in case of merges
	commit := pullRequest.Commits[0]
	noBuild := false
	restart := false
	fontello := false

	for _, file := range commit.Added {
		fmt.Println("File Added: " + file)
		if strings.HasPrefix(file, "docs/") {
			noBuild = true
		}
	}
	for _, file := range commit.Removed {
		fmt.Println("File Removed: " + file)
		if strings.HasPrefix(file, "docs/") {
			noBuild = true
		}
	}
	for _, file := range commit.Modified {
		fmt.Println("File Modified: " + file)
		if strings.HasPrefix(file, "docs/") {
			noBuild = true
		}
		if strings.Contains(file, "daemon.run") {
			restart = true
		}
		if strings.Contains(file, "fontello-config.json") {
			fontello = true
		}
	}

	if noBuild {
		fmt.Println("Doc file, no build needed")
		return
	}

	//Rebuild fontello
	if fontello {
		rebuildFontello()
	}

	//Docs hasnt been built. lets do it
	fmt.Println("Update Site running...")
	updateSite()

	if restart {
		fmt.Println("New Daemon build, restarting")

		cmdPull := exec.Command("systemctl", "restart", "blog")
		cmdPull.Dir = settings.RootPath
		cmdPull.Stdout = os.Stdout
		cmdPull.Stderr = os.Stderr
		if err := cmdPull.Run(); err != nil {
			fmt.Printf("cmdPull.Run() failed with %s\n", err)
		}
	}
}

func rebuildFontello() {
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
}

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
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
	os.RemoveAll(settings.RootPath + "/docs/author/")

	// Hugo build
	cmdBuild := exec.Command("hugo")
	cmdBuild.Dir = settings.RootPath
	out, err := cmdBuild.CombinedOutput()
	if err != nil {
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
	cmdCommit := exec.Command("git", "commit", "-m", "Webhook rebuild\n"+string(out))
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

	erryaml := yaml.Unmarshal(b, &settings)
	if erryaml != nil {
		log.Fatal(erryaml)
	}

	fmt.Printf("Starting: %v\n", settings)
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}
