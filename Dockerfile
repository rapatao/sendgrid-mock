FROM golang:1.23.2-alpine3.20 AS builder

COPY . /build/
WORKDIR /build/

RUN go mod tidy
RUN go mod vendor

RUN apk add --update --no-cache build-base

RUN make deps build

FROM alpine:3.20

COPY --from=builder /build/app /opt/sendgrid-mock

EXPOSE 3000

RUN chmod +x "/opt/sendgrid-mock"
ENTRYPOINT [ "/opt/sendgrid-mock" ]
