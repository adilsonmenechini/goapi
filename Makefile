.PHONY: help db cover tidy

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

## ----------------
## Client install
## ----------------

## make install-air - Install AIR
install-air:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s  -- -b . \
	&& sudo mv ./air /usr/local/bin \
	&& sudo chmod +x /usr/local/bin/air \
	&& air -v

## make install-swag - Install swagger
install-swag:
	go install github.com/swaggo/swag/cmd/swag@latest
	go mod tidy
	go mod vendor

## ----------------
## Backend
## ----------------

## make tidy - clean cache and update mod
tidy:
	@go clean --modcache
	@go mod tidy 
	@go mod vendor

## make db - Docker-compose UP DB
db:  
	@docker-compose up -d

## make cover - test cover
cover: 
	@echo "Running coverage"
	@go test -cover -race ./test

## make swag - Create swagger
swag:
	$$(go env GOPATH)/bin/swag init -g ./cmd/main.go -o ./api/docs

## make clean - Docker-compose Cleaning
clean: 
	@docker-compose down
	docker system prune -a --volumes
