FROM golang:latest
WORKDIR /go/src/dockerexeternalservice
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/dockerexeternalservice/app .
EXPOSE 8090
ENTRYPOINT ["./app"]