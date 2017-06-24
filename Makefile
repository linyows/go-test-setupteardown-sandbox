TEST ?= $(shell go list ./...)

ci: test integration

test:
	go test -v $(TEST)

integration:
	go test -integration -v $(TEST)

.PHONY: test integration
