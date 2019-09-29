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

fstart:
	docker-compose run react yarn start

set-up:
	go get -u ./...

finstall:
	docker-compose exec react yarn install

fclean:
	rm -rf app/node_modules app/yarn.lock

install: set-up build

enter:
	docker-compose exec go bash

full-tests:
	docker-compose exec go go test -v

single-test:
	docker-compose exec go go test -v -run
