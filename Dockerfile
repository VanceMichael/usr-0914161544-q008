FROM golang:1.24-alpine AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
RUN go test ./... && CGO_ENABLED=0 go build -o /out/server ./cmd/server

FROM alpine:3.21
RUN adduser -D -u 10001 app
USER app
COPY --from=build /out/server /server
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/server"]
