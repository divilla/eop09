FROM golang:alpine AS build
RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

COPY entityproto ../entityproto
COPY client/go.* ./
RUN go mod download
RUN go mod verify

COPY client .
RUN make build


FROM alpine:latest
RUN apk --no-cache add ca-certificates bash curl
RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /app/cmd/build/server .
COPY --from=build /app/cmd/docker/entrypoint.sh .
COPY --from=build /app/config/*.yml ./config/
COPY --from=build /app/data/* ./data/
ENTRYPOINT ["./entrypoint.sh"]
