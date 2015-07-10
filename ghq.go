package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func getGhqPath() (string, error) {
	out, err := exec.Command("ghq", "root").Output()
	if err != nil {
		return "", err
	}
	return string(out)[:len(out)-1], nil
}

func searchForRepos(rootPath string) <-chan Repository {
	repos := make(chan Repository)

	go func() {
		filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
			// skip file
			if !info.IsDir() {
				return nil
			}

			// skip directories which is not a repository
			if _, err := os.Stat(filepath.Join(path, ".git")); err != nil {
				return nil
			}

			repository := Repository{"git", path}
			repos <- repository

			return filepath.SkipDir
		})
		close(repos)
	}()

	return repos
}
