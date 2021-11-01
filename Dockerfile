# Stage 1 - the frontend build process
FROM node:16.10.0-alpine as build-stage
# ENV PORT=8000
WORKDIR /usr/src/app
COPY . ./
WORKDIR /usr/src/app/frontend
RUN yarn && \
    yarn build && \
    rm -rf frontend

# Stage 2 - the backend build process
FROM golang:1.17-alpine
WORKDIR /usr/src/app
COPY --from=build-stage /usr/src/app .
RUN apk update && \
    apk add build-base
RUN go mod download
RUN go build -o ./ ./cmd/main.go
