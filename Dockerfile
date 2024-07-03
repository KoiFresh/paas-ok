FROM golang:bookworm as printcluster-builder

COPY ./service /service
WORKDIR /service
RUN CGO_ENABLED=0 go build -o serve cmd/serve/main.go



FROM node:22-bookworm as printcluster-web-builder

COPY ./web /web
WORKDIR /web
RUN yarn install --frozen-lockfile && yarn generate



FROM debian:bullseye-slim as runner

RUN apt-get update && apt-get install -y prusa-slicer
COPY --from=printcluster-builder /service/serve /app/serve
COPY --from=printcluster-web-builder /web/dist /app/static
EXPOSE 8080
WORKDIR /app
ENTRYPOINT [ "./serve" ]
