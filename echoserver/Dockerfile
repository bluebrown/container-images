FROM golang:1.20
WORKDIR /go/src
COPY main.go ./
RUN  --mount=type=cache,target=/root/.cache/go-build \
  CGO_ENABLED=0 go build -ldflags "-w -s" -o echoserver main.go

FROM scratch
COPY --from=0 /go/src/echoserver /usr/bin/echoserver
ENTRYPOINT ["echoserver"]
