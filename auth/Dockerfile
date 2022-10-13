#First Image to compile code to binary code
FROM golang:alpine3.16 AS auth-compiler
RUN adduser -D -g '' client
WORKDIR /opt/app/
COPY go.mod go.sum ./
RUN go mod download && \
    go mod verify
COPY . .
RUN go get github.com/jackc/pgconn@v1.13.0
RUN go get github.com/jackc/pgx/v4/pgxpool@v4.17.2
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/micro-auth ./cmd/app

#Second to running binary code from first image
FROM alpine:latest
COPY --from=auth-compiler /etc/passwd /etc/passwd
COPY --from=auth-compiler --chown=client:1000 /go/bin/micro-auth /micro-auth
USER client
ENTRYPOINT ["./micro-auth"]