FROM golang:1.15-buster as builder
WORKDIR build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags '-extldflags -s -w' -o /bin/app ./

FROM scratch
COPY --from=builder /bin/app /bin/app
ENTRYPOINT ["/bin/app"]
EXPOSE 8080