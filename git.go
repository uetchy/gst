package main

import (
	"errors"
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

type RepositoryNotFoundError struct {
	TargetPath string
}

func (f RepositoryNotFoundError) Error() string {
	return "Repository not found or moved: " + f.TargetPath
}

type NoRemoteSpecifiedError struct {
	TargetPath string
}

func (f NoRemoteSpecifiedError) Error() string {
	return "No remote repository specified: " + f.TargetPath
}

type NoCommitsError struct {
	TargetPath string
}

func (f NoCommitsError) Error() string {
	return "Does not have any commits yet: " + f.TargetPath
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

// git log --branches --not --remotes
func GitLog(targetPath string) (string, error) {
	if err := os.Chdir(targetPath); err != nil {
		return "", err
	}

	out, err := exec.Command("git", "log", "--branches", "--not", "--remotes", "--oneline").CombinedOutput()
	if err != nil {
		eout := string(out)
		if strings.HasPrefix(eout, "does not have any commits yet") {
			return "", &NoCommitsError{targetPath}
		} else {
			return "", err
		}
	}

	if len(out) == 0 {
		return "", errors.New("No output")
	}

	return string(out), nil
}

func GitRemoteAdd(targetPath string, name string, url string) error {
	if err := os.Chdir(targetPath); err != nil {
		return err
	}

	_, err := exec.Command("git", "remote", "add", name, url).Output()
	if err != nil {
		return err
	}

	return nil
}

func GitRemoteSetURL(targetPath string, name string, url string) error {
	if err := os.Chdir(targetPath); err != nil {
		return err
	}

	_, err := exec.Command("git", "remote", "set-url", name, url).Output()
	if err != nil {
		return err
	}

	return nil
}

// GitPull pulls remote branch
func GitPull(targetPath string) (string, error) {
	if err := os.Chdir(targetPath); err != nil {
		return "", err
	}

	out, err := exec.Command("git", "pull").CombinedOutput()
	if err != nil {
		eout := string(out)
		if strings.HasPrefix(eout, "conq: repository does not exist.") {
			return "", &RepositoryNotFoundError{targetPath}
		} else if strings.HasPrefix(eout, "ERROR: Repository not found.") {
			return "", &RepositoryNotFoundError{targetPath}
		} else if strings.HasPrefix(eout, "fatal: No remote repository specified.") {
			return "", &NoRemoteSpecifiedError{targetPath}
		} else {
			return "", err
		}
	}

	return string(out), nil
}
