.PHONY: lint
lint:
	golangci-lint run --out-format=github-actions --enable=revive,gosec,prealloc,gocognit,bodyclose,gofmt