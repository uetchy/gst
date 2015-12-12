# gst: ghq support toolbox

[![wercker status](https://app.wercker.com/status/2715e17aa6fc187dfa5031b62df5c2e5/s "wercker status")](https://app.wercker.com/project/bykey/2715e17aa6fc187dfa5031b62df5c2e5)

## Install

```console
❯ go get github.com/uetchy/gst
```

with Homebrew:

```console
❯ brew install uetchy/gst/gst
```

## Usage
### `gst list` or `gst`

List all of repositories changed git status

```console
❯ gst
/Users/uetchy/repos/src/github.com/uetchy/cabret
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

/Users/uetchy/repos/src/github.com/uetchy/gst
-- 3 minutes ago
A  changelog.md
R  gst.go -> main.go
```

with __--short__ option:

```console
❯ gst --short
/Users/uetchy/repos/src/github.com/uetchy/ferret
/Users/uetchy/repos/src/github.com/uetchy/gst
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
/Users/uetchy/repos/src/github.com/uetchy/awesome-project
❯ gst new epic-team/awesome-project
/Users/uetchy/repos/src/github.com/epic-team/awesome-project
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

```console
❯ gst rm horrible-project
Remove? /Users/uetchy/repos/src/github.com/uetchy/horrible-project
```
