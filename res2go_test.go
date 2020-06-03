package main

import (
	"os"
	"os/exec"
	"testing"
)

func TestMainProgram(t *testing.T) {
	os.Args = []string{"res2go",
		"-package", "main",
		"-o", "checks/resource.go",
		"./tests"}
	main()
	// build check.go and launch
	cmd := exec.Command("go", "build")
	cmd.Dir = "./checks"
	err := cmd.Run()
	if err != nil {
		t.Errorf("error while building checks %v", err)
	}

	cmd = exec.Command("./checks")
	cmd.Dir = "./checks"
	err = cmd.Run()
	if err != nil {
		t.Errorf("error while running checks %v", err)
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
