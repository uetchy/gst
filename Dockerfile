FROM golang:1.11

# install ghq
RUN go get github.com/motemen/ghq
COPY test/fixture/gitconfig /root/.gitconfig

# deploy fixtures for test
RUN ghq get github/gitignore
WORKDIR /go/src/github.com/github/gitignore
RUN touch newfile
RUN rm Go.gitignore
RUN echo "*" >Node.gitignore
RUN echo "*" >committedfile
RUN git add committedfile
RUN git commit -m 'Add new file'

WORKDIR /go/src/github.com/uetchy/gst

# install deps
RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

# build gst
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/uetchy/gst
RUN gst
