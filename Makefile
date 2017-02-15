.PHONY: install run run-webserver test lint

BIN := docker run \
    -it \
    --rm \
    -v "$(PWD):/src" \
    -p 8080:8080  \
    awale-go

# Initialization =====================================================

install:
	docker build --tag=awale-go .

# Run ================================================================

run:
	$(BIN) go run src/launcher.go

run-webserver :
	$(BIN) go run src/webserver/webserver.go

# Tests ===============================================================

test:
	$(BIN) go test -v ./...

# Lint ===============================================================

lint:
	$(BIN) gofmt -w src/

# Exemple for Windows :
# docker run -it --rm -v "$(PWD):/src" awale-go go test -v ./...
# docker run -it --rm -v "$(PWD):/src" awale-go go test -v ./src/ai
# docker run -it --rm -v "$(PWD):/src" awale-go go run src/launcher.go
