FROM node:lts-alpine AS build
WORKDIR /app

RUN deluser --remove-home node

# Update distri and install wget
RUN apk update \
    && update-ca-certificates

# Install PNPM
RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile

FROM node:lts-alpine AS base
WORKDIR /app

COPY --from=build /app/node_modules .
COPY --from=build /app/package.json .
COPY --from=build /app/package-lock.json .

FROM base AS dev
ENV NODE_ENV "development"
CMD [ "pnpm", "run", "dev" ]

FROM base AS prod
ENV NODE_ENV "production"
RUN pnpm run build
