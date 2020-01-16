FROM golang:1.13.6-buster

WORKDIR /go/src/app
COPY . .
RUN go build -v
RUN go install -v

CMD ["weather"]