up:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose up -d --build

logs:
	docker-compose logs -f

fresh:
	docker-compose exec go fresh

clean:
	rm -rf api/tmp

run: clean up fresh

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
