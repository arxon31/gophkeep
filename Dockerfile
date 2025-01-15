FROM golang:1.22-bullseye as builder
WORKDIR ${GOPATH}/src/github.com/arxon31/gophkeep
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
      go build -ldflags "-w -s -X main.Build=`date +%FT%T%z`" cmd/gophkeep/gophkeep.go


FROM gcr.io/distroless/static-debian12 as prod
EXPOSE 8080
COPY --from=builder /go/src/github.com/arxon31/gophkeep/gophkeep /gophkeep
ENTRYPOINT ["./gophkeep"]