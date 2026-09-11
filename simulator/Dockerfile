FROM golang:1.25-bookworm AS build

WORKDIR /src
COPY MMA2/ ./MMA2/
COPY mma2composer/ ./mma2composer/
COPY mma2raw/ ./mma2raw/
COPY simulator/ ./simulator/
WORKDIR /src/simulator
RUN go build -trimpath -o /out/modbus-simulator-runtime ./cmd/modbus-simulator-runtime

FROM debian:bookworm-slim
COPY --from=build /out/modbus-simulator-runtime /usr/local/bin/modbus-simulator-runtime
ENTRYPOINT ["/usr/local/bin/modbus-simulator-runtime"]
