package main

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

type Repository struct {
	Type string
	Path string
}

func gitStatus(targetPath string) ([]string, error) {
	if err := os.Chdir(targetPath); err != nil {
		return nil, err
	}

	out, _ := exec.Command("git", "status", "--porcelain", "-z").Output()
	if len(out) == 0 {
		return nil, errors.New("No status changed")
	}

	statuses := strings.Split(string(out), "\x00")

	return statuses, nil
}

func getGithubUser() (string, error) {
	out, err := exec.Command("git", "config", "--get", "github.user").Output()
	if err != nil {
		return "", err
	}
	return string(out)[:len(out)-1], nil
}
