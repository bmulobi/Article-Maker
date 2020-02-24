SHELL := /bin/bash

test:
	export APP_ENV=testing; go test -v ./tests; export APP_ENV=development
install:
	cp config-example.yml config.yml
	go install -v
run:
	articlemaker
	