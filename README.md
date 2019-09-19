# Vote App

Vote App is a (GO+ReactJS) app that allows users to vote online on draft laws.

## Requirements
- GO 1.13 on your local machine
- docker latest (`docker --version`) (https://docs.docker.com/install/)
- docker-compose latest (`docker-compose --version`) (https://docs.docker.com/compose/install/)
- make (make --version)

## Installation

Run this command to run **api** and watch files

```
make run
```

Run this command to run **front** and watch files

```
makef fstart
```

## Usage

- GO API : [http://localhost](http://localhost)
- Vote App : [http://localhost:3000](http://localhost:3000)

## Useful commands

- Stop all containers `make down`
- Build images `make build`

## Contributions
- Pull last version of `develop` branch
- Create new branch that name begins with `ft-` followed by functionnality's name. Example for login : `ft-login`

## Authors / Maintainers
- Ameena AZIZ <ameena.aziz9@gmail.com>
- Kaba CONDE <contact@kabaconde.com>
- Adam Malick SOW <adam.sow97@gmail.com>
