build:
	docker build -t uetchy/gst .

run: build
	docker run --rm -v $(ghq root):/ghq -it uetchy/gst --help

push: build
	docker push uetchy/gst

build-test:
	docker build -f Dockerfile.test -t gst:test .

test: build-test
	docker run --rm -it gst:test go test github.com/uetchy/gst

