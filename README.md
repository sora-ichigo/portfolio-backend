# portfolio-api

This is a backend repository for igsr5's portfolio site.

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

## Generate(wire, sqlboiler, ...)
```sh
$ make gen
```

## Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
