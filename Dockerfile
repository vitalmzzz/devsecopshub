FROM golang:1.18.1

RUN go version
ENV GOPATH=/

COPY ./build  ./

RUN mkdir /usr/local/share/ca-certificates/my-custom-ca/

COPY ./crt/* /usr/local/share/ca-certificates/my-custom-ca/

RUN update-ca-certificates

RUN go mod download
RUN go build -o app ./cmd/main.go

EXPOSE 8080

CMD ["./app"]
