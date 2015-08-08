setup:
	godep get

run:
	godep go install
	gst

test:
	godep go test
