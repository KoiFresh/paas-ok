FROM oven/bun:latest as base

FROM base as builder
COPY . /app
WORKDIR /app
RUN bun install && bun run build

FROM node:22-alpine
EXPOSE 3000
COPY --from=builder /app/.output /app/.output
WORKDIR /app

CMD [ "./.output/server/index.mjs" ]
