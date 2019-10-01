ADDRESS := ":8877"

# Start server in development mode
.PHONY: dev
dev:
	@go run dist.go -clean-packr && \
	go run . -dev -http $(ADDRESS)

# Build a binary server ready for production at bin/server
.PHONY: dist
dist: $(packr2)
	@go run dist.go
