FROM golang:alpine AS builder

WORKDIR /testBackend

COPY . .

RUN go mod download

RUN go mod verify

FROM scratch

COPY --from=builder /go/bin/exec /go/bin/exec

EXPOSE 8080

CMD ["go/bin/exec"]