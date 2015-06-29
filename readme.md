# gst

## Install
```console
❯ go get -d github.com/uetchy/gst
```

## Usage
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
