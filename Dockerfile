# Lightweight Alpine-based
FROM golang:1.17-alpine3.15 as builder

# Build ARGS
ARG VERSION="latest-alpine-3.15"

RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -v -ldflags "-s -w -X main.programVersion=${VERSION}" -o app_executable

# Multi-stage build: only copy build result and resources
FROM alpine:3.15 as runner

LABEL original_developer="ICPAC" \
    contributor="Erick Otenyo<otenyo.erick@gmail.com>" \
    vendor="ICPAC" \
	url="https://icpac.net" \
	os.version="3.12"

RUN apk --no-cache add ca-certificates && mkdir /app
WORKDIR /app/
COPY --from=builder /app/app_executable /app/

VOLUME ["/config"]

USER 1001
EXPOSE 9000

ENTRYPOINT ["/app/app_executable"]

CMD []