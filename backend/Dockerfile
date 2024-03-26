FROM golang:1.21
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build -o simple-chat-system .
CMD ["/app/simple-chat-system"]
