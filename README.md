# Vote App

Vote App is a (Golang & VueJS) app that allows users to vote online on draft laws.

## Requirements

- GO 1.13 on your local machine
- Git
- docker latest (`docker --version`) (https://docs.docker.com/install/)
- docker-compose latest (`docker-compose --version`) (https://docs.docker.com/compose/install/)
- make (`make --version`)

## Installation

- Clone project into your `$GOPATH` through `github.com/kconde2` which is generally (`~/go/src/github.com/`)

```
git clone https://github.com/kconde2/vote-app.git
```

- Install project dependencies

```
make install
```

- Run this command to run **api** and watch files (Docker)

```
make run
```

or

- Run with local environnement (Local)

```
cd && go run main.go
```

## Usage

- GO API : [http://localhost:4000](http://localhost:4000)
- Vue App : [http://localhost:8080](http://localhost:8080)

## Postman

- API Collection link to import Vote collection : https://www.getpostman.com/collections/fc1fc5458b8afaa1a0a7
- Make sure to define `url` and `token` environnements variables

## Useful commands

- Stop all containers `make down`
- Build images `make build`
- Run only vote app `make vue`
- See Vue logs `docker-compose logs -f node`

## Usefull links

- Validations : https://github.com/qor/validations
- Directives : [click here](./DIRECTIVES.md)
- Formik use case : `http://localhost:8080/formik`

## Bugs fix

If your "go" container wont build, check build errors with

```
docker-compose up
```

If your "node" container won't build and you get the "vue-cli-service: not found" error, run

```
make f-install
```

then stop all containers and relaunch them 

```
docker-compose down &&  docker-compose up -d
```

## Run Tests (In process)

To launch all tests

```
make full-tests
```

To launch signgle test

```
docker-compose exec go go test -v -run function-name
```

## Contributions

- Pull last version of `develop` branch
- Create new branch that name begins with `ft-` followed by functionnality's name. Example for login : `ft-login`

## Authors / Maintainers

- Ameena AZIZ <ameena.aziz9@gmail.com>
- Kaba CONDE <contact@kabaconde.com>
- Adam Malick SOW <adam.sow97@gmail.com>
