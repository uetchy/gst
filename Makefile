setup:
	godep save

run:
	godep go run *.go

test:
	godep go test
