FROM golang:bookworm as builder

COPY . /service
WORKDIR /service
RUN CGO_ENABLED=0 go build -o serve cmd/serve/main.go

FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y cura-engine
COPY --from=builder /service/serve /serve
EXPOSE 8080
ENTRYPOINT [ "/serve" ]
