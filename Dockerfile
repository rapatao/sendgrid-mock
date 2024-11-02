FROM golang:1.23.2-alpine3.20 AS builder

COPY . /build/
WORKDIR /build/

RUN go mod tidy && \
    go mod vendor && \
    apk add --update --no-cache build-base && \
    make deps build

FROM alpine:3.20

COPY --from=builder /build/app /opt/sendgrid-mock
COPY --from=builder /build/web /opt/web

ENV API_KEY=""
ENV EVENT_DELIVERY_URL=""
ENV MAIL_HISTORY_DURATION=""
ENV WEB_STATIC_FILES="/opt/web"
ENV STORAGE_FILE=""

EXPOSE 3000

ENTRYPOINT [ "/opt/sendgrid-mock" ]
