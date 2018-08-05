# user-service/Dockerfile
FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/maddymanu/microservices-evan-tut/user-service

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -a -installsuffix cgo main.go repository.go handler.go database.go token_service.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/maddymanu/microservices-evan-tut/user-service .

CMD ["./user-service"]