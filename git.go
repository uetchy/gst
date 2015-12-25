package main

import (
	"errors"
	"fmt"
	"github.com/motemen/go-gitconfig"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GitConfig represents git config file
type GitConfig struct {
	UserName        string `gitconfig:"user.name"`
	UserEmail       string `gitconfig:"user.email"`
	PullRebase      bool   `gitconfig:"pull.rebase"`
	GithubUser      string `gitconfig:"github.user"`
	RemoteOriginURL string `gitconfig:"remote.origin.url"`
}

// GitConfigGet returns git config value by key
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

// GitConfigSet will set git config value by key
func GitConfigSet(targetPath string, key string, value string) (string, error) {
	out, err := exec.Command("git", "config", "--file", targetPath, "--set", key, value).Output()
	return string(out), err
}

// GitStatus returns git status of certain repository
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

// GitPull pulls remote branch
func GitPull(targetPath string) error {
	if err := os.Chdir(targetPath); err != nil {
		return err
	}

	out, err := exec.Command("git", "pull").Output()
	if err != nil {
		return err
	}

	fmt.Println(string(out))
	return nil
}
