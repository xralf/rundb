.DEFAULT_GOAL := build
MAIN := api

build:
	go build $(MAIN).go
.PHONY:build

run: build
	./$(MAIN)

clean:
	rm ./$(MAIN)

setup-db:
	psql -U postgres -d rundb -f setup.sql

requests:
	./requests.sh
