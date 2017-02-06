.PHONY: install run test lint

BIN := docker run \
    -it \
    --rm \
    -v "$(PWD):/src" \
    awale-go

# Initialization =====================================================

install:
	docker build --tag=awale-go .

# Run ================================================================

run:
	$(BIN) go run src/launcher.go

# Tests ===============================================================

test:
	$(BIN) go test -v ./...

# Lint ===============================================================

lint:
	$(BIN) gofmt -w src/
