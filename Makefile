all:
	@echo "no default"

.PHONY: test
test:
	@go test -v *.go
