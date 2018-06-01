package main

import (
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/codegangsta/cli"
)

const expected = `
uncommitted changes
 D Go.gitignore
 M Node.gitignore
?? newfile
unpushed commits
`

func TestCommandList(t *testing.T) {
	app := cli.NewApp()
	app.Flags = flagsOfList
	app.Action = doList

	var err error

	out := captureStdout(func() {
		err = app.Run([]string{"gst"})
	})

	if err != nil {
		t.Errorf("Unexpected exit code: %s", err)
	}

	line := strings.Split(strings.TrimSpace(out), "\n")
	truncated := strings.Join(line[1:len(line)-1], "\n")
	if a, e := strings.TrimSpace(truncated), strings.TrimSpace(expected); a != e {
		t.Errorf("Unexpected output\n%v", diff.LineDiff(e, a))
	}
}
