
FROM golang:1.10
COPY . /go/app/
WORKDIR /go/app/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s" -o app

FROM scratch
COPY --from=0 /go/app /
LABEL app.language=golang app.name=kanary-operator
EXPOSE 8080
ENTRYPOINT ["/app"]