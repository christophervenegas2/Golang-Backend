FROM golang:alpine AS builder

WORKDIR /testBackend

COPY . .

RUN go mod download

RUN go mod verify

RUN CGO_ENABLED=0 GOOS:linux GOARCH=amd64 go build -a -installuffix cgo -ldflags="-w -s" -o /go/bin/exec

FROM scratch

COPY --from=builder /go/bin/exec /go/bin/exec

EXPOSE 8080