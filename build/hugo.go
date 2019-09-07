package main

import (
	"fmt"
	"os"
	"os/exec"
)

func buildSite() {

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
	cmdBuild.Dir = "../"
	_, err := cmdBuild.CombinedOutput()
	if err != nil {
		fmt.Printf("cmdBuild.Run() failed with %s\n", err)
	}

}
