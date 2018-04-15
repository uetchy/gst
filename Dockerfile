FROM golang:1.10

WORKDIR /go/src/github.com/uetchy/gst

RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/uetchy/gst

RUN go get github.com/motemen/ghq
RUN echo "[ghq]\
    root = /go/src" > ~/.gitconfig