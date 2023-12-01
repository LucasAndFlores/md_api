FROM golang:1.21-alpine3.18 as builder
WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./main.go

FROM scratch
WORKDIR /
COPY --from=builder /app/api ./
COPY .env ./
EXPOSE 3000
ENTRYPOINT ["./api"]
