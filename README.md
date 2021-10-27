<div align="center">
 <h1>gst 👻</h1>
</div>

<p align="center"><img src="https://raw.githubusercontent.com/uetchy/gst/master/assets/screen.gif"/></p>

[![Actions Status: release](https://github.com/uetchy/gst/workflows/goreleaser/badge.svg)](https://github.com/uetchy/gst/actions?query=goreleaser)
[![sake.sh badge](https://sake.sh/uetchy/badge.svg)](https://sake.sh/uetchy)

**gst** is a simple toolbox that offers additional commands (`list`, `new`, `rm`, `doctor`, `update`, `fetch`) for [ghq](https://github.com/x-motemen/ghq).

See [Quick Install](https://github.com/uetchy/gst#quick-install) for the installation guide.

## Usage

### `gst list` or `gst`

List **uncommitted changes** and **unpushed commits** within all repositories.

```bash
$ gst
/Users/uetchy/Repos/src/github.com/uetchy/gst (11 minutes ago)
uncommitted changes
 M .travis.yml
 M README.md

/Users/uetchy/Repos/src/github.com/uetchy/qiita-takeout (9 hours ago)
unpushed commits
409849d returns Promise.reject
```

with `--short` or `-s` option:

```bash
$ gst --short
/Users/uetchy/Repos/src/github.com/uetchy/ferret
/Users/uetchy/Repos/src/github.com/uetchy/gst
```

### new

Create a new git repository.

Before start using `new` and `rm` command, You **must** set `github.user` in gitconfig to your GitHub username: `git config --global github.user <user>`.

```bash
$ gst new epic-project
/Users/uetchy/Repos/src/github.com/uetchy/epic-project
$ gst new epic-team/epic-project
/Users/uetchy/Repos/src/github.com/epic-team/epic-project
```

With `cd`, you can jump to the created project quickly:

```bash
cd $(gst new epic-project)
```

It's also good for having a handy alias for this workflow:

```bash
newrepo() {
  cd $(gst new ${1})
}
```

### rm

Remove a git repository. It also removes the containing directory if the deleted repository was the sole repository the parent directory had.

```bash
$ gst rm retired-project
Remove? /Users/uetchy/Repos/src/github.com/uetchy/retired-project [Y/n]
Removed /Users/uetchy/Repos/src/github.com/uetchy/retired-project
Removed /Users/uetchy/Repos/src/github.com/uetchy
```

### doctor

Health-check over all repositories.

```bash
$ gst doctor
[bitbucket.org/uetchy/scent] git remote origin has changed:
   Expected:   github.com/uetchy/google-cloud-vision-raspi-sample
   Actual:       bitbucket.org/uetchy/scent
```

### update

`git pull` to all repositories.

```bash
$ gst update
/Users/uetchy/Repos/src/github.com/uetchy/gst
Already up-to-date.
```

### fetch

`git fetch --tags --prune` to all repositories.

```bash
$ gst fetch
/Users/uetchy/Repos/src/github.com/uetchy/gst
 * [new branch]      dev        -> origin/dev
 - [deleted]         (none)     -> origin/test
 * [new tag]         v1.0.0     -> v1.0.0
```

## Quick Install

See [releases](https://github.com/uetchy/gst/releases/latest).

<!-- START mdmod {match: /v\d+?\.\d+?\.\d+?/g, replace: () => version} -->

macOS:

```bash
brew tap sake.sh/uetchy https://sake.sh/uetchy
brew install gst
```

Linux:

```bash
curl -L https://github.com/uetchy/gst/releases/download/v5.0.3/gst_linux_amd64 > /usr/local/bin/gst
chmod +x /usr/local/bin/gst
```

<!-- END mdmod -->

### Run as Docker container

You can take a glance at what `gst` do before installing the actual binary by running the containerized Docker image.

```bash
alias gst="docker run --rm -v \$(ghq root):/ghq -it uetchy/gst"
gst --help
gst list
```

### Pre-release build

macOS:

```bash
curl -L https://github.com/uetchy/gst/releases/download/pre-release/gst_darwin_amd64 > /usr/local/bin/gst
chmod +x /usr/local/bin/gst
```

Linux:

```bash
curl -L https://github.com/uetchy/gst/releases/download/pre-release/gst_linux_amd64 > /usr/local/bin/gst
chmod +x /usr/local/bin/gst
```

### Head build

```bash
go get github.com/uetchy/gst
```

## Development

PRs are welcome.

```bash
go build
./gst
```

### Test

Docker is required to run tests.

```bash
make test
```

## Contributors

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<table>
  <tr>
    <td align="center"><a href="https://uechi.io"><img src="https://avatars0.githubusercontent.com/u/431808?v=4" width="100px;" alt="Yasuaki Uechi"/><br /><sub><b>Yasuaki Uechi</b></sub></a><br /><a href="https://github.com/uetchy/gst/commits?author=uetchy" title="Code">💻</a> <a href="https://github.com/uetchy/gst/commits?author=uetchy" title="Documentation">📖</a></td>
    <td align="center"><a href="https://github.com/sinshutu"><img src="https://avatars0.githubusercontent.com/u/7629220?v=4" width="100px;" alt="NaotoSuzuki"/><br /><sub><b>NaotoSuzuki</b></sub></a><br /><a href="https://github.com/uetchy/gst/commits?author=sinshutu" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->
