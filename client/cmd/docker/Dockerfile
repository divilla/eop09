FROM golang:alpine AS build
RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

# copy module files first so that they don't need to be downloaded again if no change
COPY entityproto ../entityproto
COPY client/go.* ./
RUN go mod download
RUN go mod verify

# copy source files and build the binary
COPY client .
RUN ls -la
RUN make build


FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
RUN mkdir -p /var/log/app
WORKDIR /app/
COPY --from=build /app/cmd/build/server .
COPY --from=build /app/cmd/docker/entrypoint.sh .
COPY --from=build /app/config/*.yml ./config/
RUN ls -la
ENTRYPOINT ["./entrypoint.sh"]
