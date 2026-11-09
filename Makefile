.PHONY: build run clean

# Frontend first: the Go binary embeds its output.
build:
	cd web && npm ci && npm run build
	cd api && go build -o ../bin/api ./cmd/api

run: build
	./bin/api

clean:
	rm -rf bin api/internal/webui/dist
