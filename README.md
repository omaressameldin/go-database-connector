# lazy-panda-utils
Lazy Panda is a project aimed on managing employee time for consulting companies.

## What this is
- A monorepo of Go shared packages used by **Lazy Panda** services to upload files for **Lazy Panda** services using Go, and [Gitlab-ci](https://gitlab.com/omaressameldin/lazy-panda-utils) to run tests before merging

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