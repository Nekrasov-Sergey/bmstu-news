PWD = C:\Users\hoppl\bmstu-news
#$(shell pwd)
NAME = bmstu-news-parser

.PHONY: run
run:
	go run $(PWD)/cmd/$(NAME)/

.PHONY: first
first:
	go run $(PWD)/cmd/first-parse

.PHONY: build
build:
	go build -o bin/$(NAME) $(PWD)/cmd/$(NAME)

.PHONY: test
test:
	go test $(PWD)/... -coverprofile=cover.out

.PHONY: local
local:
	cp .dist.env .env

# Запустить миграции
.PHONY: migrate
migrate:
	go run $(PWD)/cmd/migrate
