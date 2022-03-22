FROM golang:1.14.6-alpine3.12 as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /build/conservice

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /build/conservice /app/conservice
COPY base_teste.txt ./base_teste.txt
COPY views/index.html ./views/index.html
COPY views/cpf.html ./views/cpf.html
EXPOSE 8080 8080
ENTRYPOINT ["/app/conservice"]
