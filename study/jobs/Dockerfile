FROM golang:1.17 as builder

WORKDIR /app

COPY go.mod ./

COPY cmd/ cmd/
COPY pkg/ pkg/ 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/main /app/main

ENTRYPOINT [ "./main" ]