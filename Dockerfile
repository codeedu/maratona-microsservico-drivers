FROM golang:1.13

WORKDIR /go/src/
COPY . .
RUN GOOS=linux go build driver.go -ldflags="-s -w"
CMD ["./driver"]