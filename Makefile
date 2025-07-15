GOLANG_VERSION := 1.23

test:
	docker run --rm -it -v "${PWD}:/app" -w /app golang:${GOLANG_VERSION} go test -v .

lint:
	docker run --rm -v "${PWD}:/app" -w /app golangci/golangci-lint:v2.2.2 golangci-lint run -v .
