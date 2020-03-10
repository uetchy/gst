build:
	docker build -f docker/build.dockerfile -t uetchy/gst .

run: build
	docker run --rm -v $(ghq root):/ghq -it uetchy/gst --help

push: build
	docker push uetchy/gst

build-test:
	docker build -f docker/test.dockerfile -t uetchy/gst:test .

test: build-test
	docker run --rm -it uetchy/gst:test go test github.com/uetchy/gst

