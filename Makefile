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

run: up fresh

fstart:
	docker-compose run react yarn start

finstall:
	docker-compose exec react yarn install

fclean:
	rm -rf front/node_modules front/yarn.lock
