FROM oven/bun:1 AS base
FROM nginx:1.29-alpine AS webserver

FROM base AS build

WORKDIR /app

COPY package.json bun.lock ./

RUN bun i

COPY . .

RUN bun run build

FROM webserver

WORKDIR /usr/share/nginx/html

COPY --from=build /app/dist .

EXPOSE 80