SHELL := /bin/bash

test:
	export APP_ENV=testing; go test -v ./tests; export APP_ENV=development
haha:
	echo haha   
	