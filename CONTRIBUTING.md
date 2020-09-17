# Contribution Guide

## Development Guide

```bash
git clone https://github.com/uetchy/gst.git && cd gst
go build
```

## Release Guide (Maintainers only)

```bash
VERSION=vX.X.X
make readme
git add .
git commit -m "chore: release ${VERSION}"
git tag -a "$VERSION" $VERSION
git push
```
