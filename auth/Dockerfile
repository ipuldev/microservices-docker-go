#First Image to compile code to binary code
FROM golang:alpine AS auth-compiler
RUN adduser -D -g '' client
WORKDIR /opt/app/
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/micro-auth ./cmd/app

#Second to running binary code from first image
FROM alpine:latest
COPY --from=auth-compiler /etc/passwd /etc/passwd
COPY --from=auth-compiler --chown=client:1000 /go/bin/micro-auth /micro-auth
USER client
ENTRYPOINT ["./micro-auth"]