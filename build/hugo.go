package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func buildSite() error {

	dir, errwd := os.Getwd()
	if errwd != nil {
		return errwd
	}

	// Delete folders to regenerate
	if err := os.RemoveAll("../docs/page/"); err != nil {
		return err
	}
	if err := os.RemoveAll("../docs/categories/"); err != nil {
		return err
	}
	if err := os.RemoveAll("../docs/post/"); err != nil {
		return err
	}
	if err := os.RemoveAll("../docs/samples/"); err != nil {
		return err
	}
	if err := os.RemoveAll("../docs/tags/"); err != nil {
		return err
	}
	if err := os.RemoveAll("../docs/author/"); err != nil {
		return err
	}
	if err := os.RemoveAll("../docs/admin/"); err != nil {
		return err
	}

	// Hugo build
	cmdBuild := exec.Command("hugo")
	cmdBuild.Dir, _ = filepath.Abs(dir + "/../")
	output, err := cmdBuild.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Println(string(output))
	return nil
}
