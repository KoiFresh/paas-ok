FROM oven/bun:latest as base

FROM base as builder
COPY . /app
WORKDIR /app
RUN bun install && bun run build

FROM base as runner
EXPOSE 3000
COPY --from=builder /app/.output /app/.output
WORKDIR /app

ENTRYPOINT [ "bun" ]
CMD [ "run", "./.output/server/index.mjs" ]
