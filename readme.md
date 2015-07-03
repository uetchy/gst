# gst

[![wercker status](https://app.wercker.com/status/2715e17aa6fc187dfa5031b62df5c2e5/s "wercker status")](https://app.wercker.com/project/bykey/2715e17aa6fc187dfa5031b62df5c2e5)

## Install

```console
❯ go get -d github.com/uetchy/gst
```

on OS X:

```console
❯ brew install uetchy/gst/gst
```

## Usage
### list

```console
❯ gst
/Users/uetchy/repos/src/github.com/uetchy/ferret
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
A  changelog.md
M  gst.go
```

with __--short__ option:

```console
❯ gst --short
/Users/uetchy/repos/src/github.com/uetchy/ferret
/Users/uetchy/repos/src/github.com/uetchy/gst
```

You can use `peco` for pipeline processing as:

```
❯ cd $(gst --short | peco)
```

### new

Before start using 'new' and 'rm' command, You __must__ set 'github.user' to .gitconfig with `git config --global github.user <user>`.

```console
❯ gst new awesome-project
/Users/uetchy/repos/src/github.com/uetchy/awesome-project
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
❯ gst rm awesome-project
```
