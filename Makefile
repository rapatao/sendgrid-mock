all: deps build

deps:
	go mod tidy
	go mod vendor

run:
	go run ./...

build:
	CGO_ENABLED=1 go build -mod vendor -a -o ./app ./cmd/...

docker-build:
	docker build -t sendgrid-mock:develop .

docker-run:
	docker run -p 3000:3000 --rm sendgrid-mock:develop

docker: docker-build docker-run
