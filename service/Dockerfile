FROM golang:bookworm as builder

COPY . /service
WORKDIR /service
RUN CGO_ENABLED=0 go build -o serve cmd/serve/main.go



FROM ubuntu:24.04

COPY --from=builder /service/resources /resources
COPY --from=builder /service/serve /serve
RUN apt-get update && \
	DEBIAN_FRONTEND=noninteractive apt-get install -y prusa-slicer && \
	rm -rf /var/lib/apt/lists/*

EXPOSE 8080
ENTRYPOINT [ "/serve" ]
