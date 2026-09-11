FROM golang:1.25-bookworm AS build

WORKDIR /src
COPY MMA2/ ./MMA2/
COPY mma2composer/ ./mma2composer/
COPY mma2raw/ ./mma2raw/
COPY simulator/ ./simulator/
COPY replicator/ ./replicator/
WORKDIR /src/replicator
RUN go build -trimpath -o /out/modbus-replicator-runtime ./cmd/modbus-replicator-runtime

FROM debian:bookworm-slim
COPY --from=build /out/modbus-replicator-runtime /usr/local/bin/modbus-replicator-runtime
ENTRYPOINT ["/usr/local/bin/modbus-replicator-runtime"]
