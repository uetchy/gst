package main

import (
	"errors"
	"github.com/motemen/go-gitconfig"
	"os"
	"os/exec"
	"strings"
	"path/filepath"
)

type GitConfig struct {
	UserName        string `gitconfig:"user.name"`
	UserEmail       string `gitconfig:"user.email"`
	PullRebase      bool   `gitconfig:"pull.rebase"`
	GithubUser      string `gitconfig:"github.user"`
	RemoteOriginURL string `gitconfig:"remote.origin.url"`
}

func GitConfigGet(targetPath string, key string) (string, error) {
	var configFile gitconfig.Config
	switch targetPath {
	case "global":
		configFile = gitconfig.Global
	case "local":
		configFile = gitconfig.Local
	default:
		configPath := filepath.Join(targetPath, ".git/config")
		configFile = gitconfig.File(configPath)
	}
	result, err := configFile.GetString(key)
	if err != nil {
		return "", err
	}

	return result, nil
}

func GitConfigSet(targetPath string, key string, value string) (string, error) {
	out, err := exec.Command("git", "config", "--file", targetPath, "--set", key, value).Output()
	return string(out), err
}

func GitStatus(targetPath string) ([]string, error) {
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
