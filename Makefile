generate:
	go generate ./...
.PHONY: generate

check_generated: generate
	git diff --exit-code
.PHONY: check_generated