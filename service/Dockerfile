FROM golang:bookworm as builder

COPY . /app
WORKDIR /app
RUN go build -o runner cmd/runner/main.go


FROM scratch as runner

COPY --from=builder /app/runner /app/runner
ENTRYPOINT [ "/app/runner" ]
