FROM node:22-bookworm as base

FROM base as builder
COPY . /app
WORKDIR /app
RUN yarn install --frozen-lockfile && yarn build

FROM base as runner
EXPOSE 3000
COPY --from=builder /app/.output /app/.output
WORKDIR /app

ENTRYPOINT [ "node" ]
CMD [ ".output/server/index.mjs" ]
