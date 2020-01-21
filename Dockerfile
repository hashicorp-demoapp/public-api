FROM alpine:latest

RUN addgroup api && \
  adduser -D -G api api

RUN mkdir /app
COPY ./bin/public-api /app/public-api

USER api

ENTRYPOINT [ "/app/public-api" ]