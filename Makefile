.PHONY: build

build:
	CGO_ENABLED=0 go build -o stemasm cmd/stemasm/stemasm.go
