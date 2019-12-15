# Go Database Connector
This is a project that creates a go database connector to any kind of database.


## What this is
- A Go database connector to implement adapters for dadtabase types.
- Right now it only implements a connector for firebase with the possibility of other extensions later on.
- [Gitlab-ci](https://gitlab.com/omaressameldin/go-database-connector) is used to run tests on merging

## How to run
- make sure you have **docker version: 18.x+** installed
- run `docker-compose up --build` to run all tests locally

## Packages included

- [database](./app/pkg/database)
- [firebase](./app/pkg/firebase)

## Technologies used
- Golang
- firebase
- Gitlab-ci
- Docker
- Docker-compose