# Docker compose commands
local:
	echo "Starting local environment with docker"
	docker-compose -f docker-compose.local.yml up --build

# Main
run:
	go run ./cmd/main.go

build:
	go build ./cmd/main.go

test:
	go test -v -cover ./...

# Modules support
deps-reset:
	git checkout -- go.mod
	go mod tidy
	go mod vendor

tidy:
	go mod tidy
	go mod vendor

deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy
	go mod vendor

deps-clean-cache:
	go clean -modcache


# Docker support
FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)