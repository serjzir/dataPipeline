FROM golang:latest as builder

WORKDIR /app

COPY . ./
RUN ls -la
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo cmd/main/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/config.yml .

CMD ["./main"]
