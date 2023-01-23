INTERNAL = ./internal/...

.PHONY: all
all:
	make clear
	make run-linter
	make build
	./main.out

.PHONY: gen
gen:
	go generate $(INTERNAL)

.PHONY: run-linter
run-linter:
	golangci-lint run $(INTERNAL) --config=./linter_config/config.yml
	go fmt $(INTERNAL)

.PHONY: build
build:
	go build -o main.out cmd/main.go

.PHONY: clear
clear:
	rm -f main.out

.PHONY: docker-build
docker-build:
	docker build --no-cache -t park .

.PHONY: docker-run
docker-run:
	docker run --rm -p 5000:5000 -p 5432:5432 --name park -t park
