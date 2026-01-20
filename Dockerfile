FROM golang:1.21-alpine AS build

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/ayma ./cmd/server

FROM alpine:3.19
RUN adduser -D -g '' appuser && apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /out/ayma /app/ayma
COPY web /app/web

ENV ADDR=:8080
EXPOSE 8080

USER appuser
CMD ["/app/ayma"]
