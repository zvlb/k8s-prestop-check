##
## Build
##
FROM golang:1.22 AS builder

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH} go build -a -o bin/app main.go 

##
## Deploy
##
FROM alpine:3.19

USER 65532:65532
WORKDIR /opt

COPY --from=builder /app/bin/app .

CMD ["/opt/app"]
