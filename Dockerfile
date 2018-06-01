FROM golang:1.10


# install ghq
RUN go get github.com/motemen/ghq
COPY test/fixture/gitconfig /root/.gitconfig

# deploy fixtures for test
RUN ghq get github/gitignore && \
    cd /go/src/github.com/github/gitignore && \
    touch newfile && \
    rm Go.gitignore && \
    echo "*" > Node.gitignore && \
    echo "*" > committedfile && \
    git add committedfile && \
    git commit -m 'Add new file'

WORKDIR /go/src/github.com/uetchy/gst

# install deps
RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

# build gst
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/uetchy/gst