# gst: ghq support toolbox

[![Build Status](https://travis-ci.org/uetchy/gst.svg)](https://travis-ci.org/uetchy/gst)

## Install

```console
❯ go get github.com/uetchy/gst
```

## Usage
### `gst list` or `gst`

List all of repositories changed git status

```console
❯ gst
/Users/uetchy/Repos/src/github.com/uetchy/cabret
-- 4 hours ago
A  .eslintrc
M  .gitignore
A  app/ArticleList.jsx
A  app/Header.jsx
A  app/app.jsx
M  index.css
M  index.html
M  index.js
M  package.json

/Users/uetchy/Repos/src/github.com/uetchy/gst
-- 3 minutes ago
A  changelog.md
R  gst.go -> main.go
```

with __--short__ option:

```console
❯ gst --short
/Users/uetchy/Repos/src/github.com/uetchy/ferret
/Users/uetchy/Repos/src/github.com/uetchy/gst
```

You can also use `peco` for pipeline processing as:

```
❯ cd $(gst --short | peco)
```

### new

Create git repository.

Before start using 'new' and 'rm' command, You __must__ set 'github.user' to .gitconfig with `git config --global github.user <user>`.

```console
❯ gst new awesome-project
/Users/uetchy/Repos/src/github.com/uetchy/awesome-project
❯ gst new epic-team/awesome-project
/Users/uetchy/Repos/src/github.com/epic-team/awesome-project
```

with `cd`, You can jump to created project:

```console
❯ cd $(gst new awesome-project)
```

or with declare function:

```zsh
newrepo() {
  cd $(gst new ${1})
}
```

### rm

Remove git repository.

```console
❯ gst rm horrible-project
Remove? /Users/uetchy/Repos/src/github.com/uetchy/horrible-project
```

### doctor

Health-check for repositories.

```console
❯ gst doctor
[bitbucket.org/uetchy/scent] git remote origin has changed:
   Expected:   github.com/uetchy/google-cloud-vision-raspi-sample
   Actual:	   bitbucket.org/uetchy/scent
```

### update

`git pull` for all repositories.

```console
❯ gst update
/Users/uetchy/Repos/src/github.com/uetchy/gst
Already up-to-date.
```
