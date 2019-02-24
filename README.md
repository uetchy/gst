<div align="center">
 <h1>gst 👻</h1>
</div>

[![Build Status](https://travis-ci.org/uetchy/gst.svg)](https://travis-ci.org/uetchy/gst)

**gst** is a auxiliary toolbox for [motemen/ghq](https://github.com/motemen/ghq).

This tool provides additional commands (`list`, `new`, `rm`, `doctor`, `update`, `fetch`) over ghq enabled environment.

![](http://uechi-public.s3.amazonaws.com/github/gst.gif)

See [Quick Install](https://github.com/uetchy/gst#quick-install) for the installation guide.

# Usage

## `gst list` or `gst`

List **uncommitted changes** and **unpushed commits** within all repositories.

```
$ gst
/Users/uetchy/Repos/src/github.com/uetchy/gst (11 minutes ago)
uncommitted changes
 M .travis.yml
 M README.md

/Users/uetchy/Repos/src/github.com/uetchy/qiita-takeout (9 hours ago)
unpushed commits
409849d returns Promise.reject
```

with **--short** option:

```
$ gst --short
/Users/uetchy/Repos/src/github.com/uetchy/ferret
/Users/uetchy/Repos/src/github.com/uetchy/gst
```

You can also use it with `peco` for pipeline processing as:

```
cd $(gst --short | peco)
```

## new

Create a new git repository.

Before start using 'new' and 'rm' command, You **must** set 'github.user' to .gitconfig with `git config --global github.user <user>`.

```
$ gst new epic-project
/Users/uetchy/Repos/src/github.com/uetchy/epic-project
$ gst new epic-team/epic-project
/Users/uetchy/Repos/src/github.com/epic-team/epic-project
```

with `cd`, You can jump to created project:

```
$ cd $(gst new epic-project)
```

or with declare function:

```zsh
newrepo() {
  cd $(gst new ${1})
}
```

## rm

Remove a git repository.

```
$ gst rm retired-project
Remove? /Users/uetchy/Repos/src/github.com/uetchy/retired-project
```

## doctor

Health-check over all repositories.

```
$ gst doctor
[bitbucket.org/uetchy/scent] git remote origin has changed:
   Expected:   github.com/uetchy/google-cloud-vision-raspi-sample
   Actual:       bitbucket.org/uetchy/scent
```

## update

`git pull` to all repositories.

```
$ gst update
/Users/uetchy/Repos/src/github.com/uetchy/gst
Already up-to-date.
```

## fetch

`git fetch --tags --prune` to all repositories.

```
$ gst fetch
/Users/uetchy/Repos/src/github.com/uetchy/gst
 * [new branch]      dev        -> origin/dev
 - [deleted]         (none)     -> origin/test
 * [new tag]         v1.0.0     -> v1.0.0
```

# Quick Install

See [releases](https://github.com/uetchy/gst/releases/latest).

macOS:

```
curl -L https://github.com/uetchy/gst/releases/download/v5.0.1/gst_darwin_amd64 > /usr/local/bin/gst
chmod +x /usr/local/bin/gst
```

Linux:

```
curl -L https://github.com/uetchy/gst/releases/download/v5.0.1/gst_linux_amd64 > /usr/local/bin/gst
chmod +x /usr/local/bin/gst
```

## Run as Docker container

You can take a glance at what `gst` do by running Docker container, before installing the actual binary.

```bash
alias gst="docker run --rm -v \$(ghq root):/ghq -it uetchy/gst"
gst --help
gst list
```

## Pre-release build

macOS:

```
curl -L https://github.com/uetchy/gst/releases/download/pre-release/gst_darwin_amd64 > /usr/local/bin/gst
chmod +x /usr/local/bin/gst
```

Linux:

```
curl -L https://github.com/uetchy/gst/releases/download/pre-release/gst_linux_amd64 > /usr/local/bin/gst
chmod +x /usr/local/bin/gst
```

## Head build

```
go get github.com/uetchy/gst
```

# Contribution

PRs are welcome.

## Test

Docker is required to run tests.

```bash
make test
```

## Contributors

- Yasuaki Uechi
- NaotoSuzuki
