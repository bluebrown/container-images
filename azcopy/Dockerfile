FROM golang:1.20-alpine3.17
RUN apk --no-cache add git
WORKDIR /workspace
RUN git clone --branch 10.18.1 --single-branch https://github.com/Azure/azure-storage-azcopy .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o "dist/azcopy"

FROM gcr.io/distroless/static:nonroot
COPY --from=0 /workspace/dist/azcopy /usr/bin/azcopy
ENTRYPOINT [ "azcopy" ]

