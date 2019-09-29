up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose up -d --build

logs:
	docker-compose logs -f

go-logs:
	docker-compose logs -f go

fresh:
	docker-compose exec go fresh

clean:
	rm -rf api/tmp

run: go-logs

set-up:
	go get -u ./...

install: set-up build

enter:
	docker-compose exec go bash

full-tests:
	docker-compose exec go go test -v
