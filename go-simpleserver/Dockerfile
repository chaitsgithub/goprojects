FROM golang:latest
WORKDIR /go/src/app
COPY . .
RUN go build -o server .
EXPOSE 8080
CMD ["./server"]