FROM ghcr.io/wormhole-foundation/aptos:0.4.0@sha256:0d88ac1bc733fb0ed502f6352a1453a14f31404e7f7f4c1b978c97c62cefbff4 as aptos

# Support additional root CAs
COPY cert.pem* /certs/
RUN if [ -e /certs/cert.pem ]; then cp /certs/cert.pem /etc/ssl/certs/ca-certificates.crt; fi

COPY start_node.sh /tmp/start_node.sh

FROM aptos AS tests

WORKDIR /tmp

COPY wormhole/ wormhole
COPY token_bridge/ token_bridge
COPY deployer/ deployer
COPY coin/ coin
COPY examples/ examples
COPY Makefile Makefile

RUN --mount=type=cache,target=/root/.move,id=move_cache \
    make test
