FROM golang
WORKDIR /workspace
COPY main.go ./
ARG CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -a -installsuffix static -o echoserver main.go

FROM gcr.io/distroless/static:nonroot
COPY --from=0 /workspace/echoserver /usr/bin/echoserver
ENTRYPOINT ["echoserver"]
