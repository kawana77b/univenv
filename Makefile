test:
	go test -v ./...

.PHONY: test

TEST_FILE := ./tests/config.yml

run:
	go run main.go fish -f $(TEST_FILE)

run-bash:
	go run main.go bash -f $(TEST_FILE)

run-fish:
	go run main.go fish -f $(TEST_FILE)

run-pwsh:
	go run main.go pwsh -f $(TEST_FILE)

.PHONY: run run-bash run-fish run-pwsh

help-root:
	go run main.go --help

.PHONY: help-init

cp:
	mkdir -p ~/.config/univenv && touch ~/.config/univenv/config.yml
	cp ./tests/config.yml ~/.config/univenv

.PHONY: cp

release:
	goreleaser release --snapshot --clean

.PHONY: release

clean:
	@rm -rf dist

.PHONY: clean