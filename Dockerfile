FROM alpine:latest as base

RUN addgroup api && \
  adduser -D -G api api

RUN mkdir /app

# Setup AMD64
FROM base as image-amd64
COPY ./bin/amd64/public-api /app/public-api
RUN chmod +x /app/public-api

# Setup ARM64
FROM base as image-arm64
COPY ./bin/arm64/public-api /app/public-api
RUN chmod +x /app/public-api

FROM image-${TARGETARCH}

RUN echo "Running on $BUILDPLATFORM platform, building for $TARGETPLATFORM"

USER api

ENTRYPOINT [ "/app/public-api" ]