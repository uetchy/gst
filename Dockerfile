FROM golang:1.11 as BUILD

# install ghq
RUN go get github.com/motemen/ghq

# install deps
WORKDIR /go/src/github.com/uetchy/gst
RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v -vendor-only

# build gst
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go install -ldflags="-w -s" -v github.com/uetchy/gst

# copy binaries from build step
FROM alpine
ENV GHQ_ROOT /ghq
ENV PATH /go/bin:$PATH
RUN apk add git
COPY --from=BUILD /go/bin/ghq /go/bin/ghq
COPY --from=BUILD /go/bin/gst /go/bin/gst
ENTRYPOINT ["gst"]
