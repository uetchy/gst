package main

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Repository represents git repository
type Repository struct {
	Type    string
	Path    string
	ModTime time.Time
}

// Repositories contains array of Repository
type Repositories []Repository

// Len return number of repositories
func (r Repositories) Len() int {
	return len(r)
}

// Swap repository
func (r Repositories) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// RepositoriesByModTime is wrapper of sort algorithm for order by mod time
type RepositoriesByModTime struct {
	Repositories
}

// Less sort array by mod time
func (bmt RepositoriesByModTime) Less(i, j int) bool {
	return bmt.Repositories[i].ModTime.Before(bmt.Repositories[j].ModTime)
}

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

			repository := Repository{
				Type:    "git",
				Path:    path,
				ModTime: info.ModTime(),
			}
			repos <- repository

			return filepath.SkipDir
		})
		close(repos)
	}()

	return repos
}

var hasSchemePattern = regexp.MustCompile("^[^:]+://")
var scpLikeURLPattern = regexp.MustCompile("^([^@]+@)?([^:]+):/?(.+)$")

func formatURL(ref string) (*url.URL, error) {
	if !hasSchemePattern.MatchString(ref) && scpLikeURLPattern.MatchString(ref) {
		matched := scpLikeUrlPattern.FindStringSubmatch(ref)
		user := matched[1]
		host := matched[2]
		path := matched[3]

		ref = fmt.Sprintf("ssh://%s%s/%s", user, host, path)
	}

	url, err := url.Parse(ref)
	if err != nil {
		return url, err
	}

	if !url.IsAbs() {
		if !strings.Contains(url.Path, "/") {
			url.Path = url.Path + "/" + url.Path
		}
		url.Scheme = "https"
		url.Host = "github.com"
		if url.Path[0] != '/' {
			url.Path = "/" + url.Path
		}
	}

	return url, nil
}

func compileTargetPathFromURL(query string) string {
	source, _ := formatURL(query)
	encodedPath := strings.TrimSuffix(source.Path, ".git")
	ghqPath := filepath.Join(source.Host, encodedPath)
	return ghqPath
}

func compileTargetPath(query string) string {
	ghqPath, err := getGhqPath()
	if err != nil {
		fmt.Println("You must setup 'ghq' command")
		os.Exit(1)
	}

	re, _ := regexp.Compile("^(?:(?:(.+?)/)?(.+?)/)?(.+)$")
	res := re.FindStringSubmatch(query)

	targetHost := res[1]
	targetUser := res[2]
	targetPath := res[3]

	if res[1] == "" {
		targetHost = "github.com"
	}

	if res[2] == "" {
		targetUser, err = GitConfigGet("global", "github.user")
		if err != nil {
			fmt.Println("You must set github.user first")
			fmt.Println("> git config --global github.user <name>")
			os.Exit(1)
		}
	}

	return filepath.Join(ghqPath, targetHost, targetUser, targetPath)
}
