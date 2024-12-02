.DEFAULT_GOAL := build

projects := day1

.PHONY: lint
lint:
	@$(foreach project,$(projects),golangci-lint run ./cmd/$(project)/... ;)


.PHONY: build
build: lint
	@$(foreach project,$(projects),go build -o ./target/$(project) ./cmd/$(project)/main.go ;)

.PHONY: run/day1
run/day1:
	@./target/day1 -i "cmd/day1/input1.txt"

.PHONY: install/tools
install/tools:
	@brew install golangci-lint
	@brew upgrade golangci-lint