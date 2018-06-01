test: build
	docker run --rm -it gst go test github.com/uetchy/gst

build: 
	docker build -t gst .
