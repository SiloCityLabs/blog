package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func buildSite() {

	dir, errwd := os.Getwd()
	if errwd != nil {
		log.Fatal(errwd)
	}

	// Delete folders to regenerate
	os.RemoveAll("../docs/page/")
	os.RemoveAll("../docs/categories/")
	os.RemoveAll("../docs/post/")
	os.RemoveAll("../docs/samples/")
	os.RemoveAll("../docs/tags/")
	os.RemoveAll("../docs/author/")
	os.RemoveAll("../docs/admin/")

	// Hugo build
	cmdBuild := exec.Command("hugo")
	cmdBuild.Dir, _ = filepath.Abs(dir + "/../")
	output, err := cmdBuild.CombinedOutput()
	if err != nil {
		log.Fatalf("cmdBuild.Run() failed with %s\n", err.Error())
	}

	fmt.Println(output)
}
