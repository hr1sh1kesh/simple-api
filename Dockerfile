FROM golang:1.11.4 as builder
WORKDIR /go/src/github.com/hr1sh1kesh/simple-api
COPY go.mod go.sum ./
RUN export GO111MODULE=on && go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/simple-api.o .

FROM alpine
WORKDIR /root/
EXPOSE 10000
COPY --from=builder /go/src/github.com/hr1sh1kesh/simple-api/bin .
ENTRYPOINT [ "./simple-api.o" ]
