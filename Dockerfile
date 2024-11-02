FROM golang:1.23.2-alpine3.20 AS builder

COPY . /build/
WORKDIR /build/

RUN go mod tidy
RUN go mod vendor

RUN apk add --update --no-cache build-base

RUN make deps build

FROM alpine:3.20

COPY --from=builder /build/app /opt/sendgrid-mock
COPY --from=builder /build/web /opt/web

ENV API_KEY=""
ENV EVENT_DELIVERY_URL=""
ENV MAIL_HISTORY_DURATION=""
ENV WEB_STATIC_FILES="/opt/web"
ENV STORAGE_DIR="/tmp/"

EXPOSE 3000

RUN chmod +x "/opt/sendgrid-mock"
ENTRYPOINT [ "/opt/sendgrid-mock" ]
