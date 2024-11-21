FROM golang:1.23-bookworm

RUN go install golang.org/x/pkgsite/cmd/pkgsite@latest

# GOPATH=go, but $GOPATH does not work here...
CMD ["/go/bin/pkgsite"]
