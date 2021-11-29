FROM golang:latest as builder
COPY go.mod go.sum /go/src/github.com/cecobask/redhat-coding-challenge/
WORKDIR /go/src/github.com/cecobask/redhat-coding-challenge
RUN go mod download
COPY . /go/src/github.com/cecobask/redhat-coding-challenge
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/redhat-coding-challenge github.com/cecobask/redhat-coding-challenge

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/cecobask/redhat-coding-challenge/build/redhat-coding-challenge /usr/bin/redhat-coding-challenge
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/redhat-coding-challenge"]