#First Image to compile code to binary code
FROM golang:alpine AS product-compiler
RUN adduser -D -g '' client
WORKDIR /opt/app/
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/micro-product ./cmd/app

#Second to running binary code from first image
FROM alpine:3.14.1
COPY --from=product-compiler /etc/passwd /etc/passwd
COPY --from=product-compiler --chown=client:1000 /go/bin/micro-product /micro-product
USER client
ENTRYPOINT ["./micro-product"]