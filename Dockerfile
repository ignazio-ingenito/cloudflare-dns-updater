# Generate template
FROM ghcr.io/a-h/templ:latest AS templ
WORKDIR /app
COPY --chown=65532:65532 views/ views/

# Build
FROM golang:alpine AS build

RUN apk add --update alpine-sdk

WORKDIR /app

COPY go.mod go.sum ./
COPY --chown=65532:65532 api/ api/
COPY --chown=65532:65532 cron/ cron/
COPY --chown=65532:65532 db/ db/
COPY --chown=65532:65532 icons/ icons/
COPY --chown=65532:65532 models/ models/
COPY --chown=65532:65532 types/ types/
COPY --chown=65532:65532 web/ web/
COPY --chown=65532:65532 *.go .

COPY --from=templ /app/views /app/views

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /app/build/dns-updater .

FROM alpine AS prod
RUN apk add --no-cache tzdata

WORKDIR /opt/flying/bin

COPY --from=build /app/build/dns-updater /opt/flying/bin/dns-updater

EXPOSE 80
CMD [ "./dns-updater" ]