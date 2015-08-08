package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func gitStatus(targetPath string) ([]string, error) {
	if err := os.Chdir(targetPath); err != nil {
		return nil, err
	}

	out, _ := exec.Command("git", "status", "--porcelain").Output()
	if len(out) == 0 {
		return nil, errors.New("No status changed")
	}

	statuses := strings.Split(string(out), "\n")

	return statuses, nil
}

func getGithubUser() (string, error) {
	out, err := exec.Command("git", "config", "--get", "github.user").Output()
	if err != nil {
		return "", err
	}
	return string(out)[:len(out)-1], nil
}
