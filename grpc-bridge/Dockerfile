# Build stage
FROM golang:1.22.4-bullseye as BuildStage

WORKDIR /app

COPY .. .

EXPOSE 9091

RUN go build ./cmd/grpc_bridge.go

#Deploy stage
ARG BUILD_FROM=ghcr.io/hassio-addons/ubuntu-base:10.0.1

FROM ${BUILD_FROM}

WORKDIR /app

COPY --from=BuildStage /app/grpc_bridge /grpc_bridge

# Build arguments
ARG BUILD_ARCH
ARG BUILD_DATE
ARG BUILD_DESCRIPTION
ARG BUILD_NAME
ARG BUILD_REF
ARG BUILD_REPOSITORY
ARG BUILD_VERSION

EXPOSE 9091

ENTRYPOINT [ "/grpc_bridge" ]

# Labels
LABEL \
    io.hass.name="${BUILD_NAME}" \
    io.hass.description="${BUILD_DESCRIPTION}" \
    io.hass.arch="${BUILD_ARCH}" \
    io.hass.type="addon" \
    io.hass.version=${BUILD_VERSION} \
    maintainer="Rick Moonen <rick@moonen.network>" \
    org.opencontainers.image.title="${BUILD_NAME}" \
    org.opencontainers.image.description="${BUILD_DESCRIPTION}" \
    org.opencontainers.image.vendor="Moonen Home Assistant Services" \
    org.opencontainers.image.authors="Rick Moonen <rick@moonen.network>" \
    org.opencontainers.image.licenses="Apache2" \
    org.opencontainers.image.url="https://github.com/${BUILD_REPOSITORY}" \
    org.opencontainers.image.source="https://github.com/${BUILD_REPOSITORY}" \
    org.opencontainers.image.documentation="https://github.com/${BUILD_REPOSITORY}/blob/main/README.md" \
    org.opencontainers.image.created=${BUILD_DATE} \
    org.opencontainers.image.revision=${BUILD_REF} \
    org.opencontainers.image.version=${BUILD_VERSION}