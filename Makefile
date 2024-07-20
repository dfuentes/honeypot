.PHONY: build
build:
	CGO_ENABLED=0 go build

.PHONY: docker-build
docker-build: build
	docker buildx build -t dfuentes/honeypot:latest .

.PHONY: run
run: build
	./honeypot

.PHONY: docker-run
docker-run: docker-build
	docker run -it -p 8080:8080 dfuentes/honeypot
