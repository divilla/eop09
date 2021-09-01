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
COPY server/go.* ./
RUN go mod download
RUN go mod verify

COPY server .
RUN make build


FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /app/cmd/build/server .
COPY --from=build /app/cmd/docker/entrypoint.sh .
COPY --from=build /app/config/*.yml ./config/
ENTRYPOINT ["./entrypoint.sh"]
