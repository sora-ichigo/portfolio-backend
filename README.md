# portfolio-backend
[![Build](https://github.com/igsr5/portfolio-backend/actions/workflows/build.yml/badge.svg)](https://github.com/igsr5/portfolio-server-api/actions/workflows/build.yml)
[![Deploy](https://github.com/igsr5/portfolio-backend/actions/workflows/deploy.yml/badge.svg)](https://github.com/igsr5/portfolio-server-api/actions/workflows/deploy.yml)
[![Reviewdog](https://github.com/igsr5/portfolio-backend/actions/workflows/reviewdog.yml/badge.svg)](https://github.com/igsr5/portfolio-server-api/actions/workflows/reviewdog.yml)
[![Test](https://github.com/igsr5/portfolio-backend/actions/workflows/test.yml/badge.svg)](https://github.com/igsr5/portfolio-server-api/actions/workflows/test.yml)
![](https://img.shields.io/badge/license-MIT-green)

This is a backend repository for igsr5's portfolio site.


Frontend Repo - https://github.com/igsr5/portfolio-front  
Schema Repo(Protocol Buffers) - https://github.com/igsr5/portfolio-proto

## Development

## Init
```sh
$ docker compose up -d
$ make migrate
```

## Local development
```sh
$ make sam-local
```

### Migration
first
```
$ docker compose up -d
```
second, in docker container
```sh
# create migration file
$ make migrate-create T=[FILENAME]
# migrate
$ make migrate
# rollback(1)
$ make rollback
# rollback(2~)
$ make rollback STEP=2
# fix version
$ make migrate-force VERSION=[VERSION]
```

### Test
```sh
$ make test
```

### Generate(wire, sqlboiler, ...)
```sh
$ make gen
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
