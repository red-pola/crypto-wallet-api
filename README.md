# CryptoWallet Platform API

This platform is designed for people who has crypto and want to help others by investing/donating them.
At first stage, only Stellar (XLM) is supported.

## Development

### Using Docker

#### Get Docker

Get [Docker Community Edition](https://www.docker.com/community-edition) for your platform *(Linux, Mac, Windows)*.

#### Spin up the API server and Database

```
$ docker-compose up -d
```

Run this command to verify:

```
$ curl -i localhost:8080/status
```

It should output something like this:

```
HTTP/1.1 200 OK
Date: Wed, 13 Jun 2018 05:45:36 GMT
ontent-Length: 0
```

#### Stop the server

```
$ docker-compose down
```

#### Rebuild image

```
$ docker-compose up -d --build
```

#### Access to running containers

```
$ docker-compose exec api sh
$ docker-compose exec mongo-db mongo -u root -p password --authenticationDatabase "admin"
```

Or

```
$ docker container exec -it crypto-wallet-api_api_1 sh
$ docker container exec -it crypto-wallet-api_mongo-db_1 mongo -u root -p password --authenticationDatabase "admin"
```

#### View logs

```
$ docker-compose logs api
$ docker-compose logs mongo-db
```

Or

```
$ docker container logs crypto-wallet-api_api_1
$ docker container logs crypto-wallet-api_mongo-db_1
```

### Using your own local environment

#### Install Go

#### Setup Go Workspace

#### Install MongoDB

#### Set environment variables

#### Build and Run

## API

### Projects

| Resource                      | Method | Description                       |
|:------------------------------|:-------|:----------------------------------|
| `/api/v1/projects`            | GET    | Get all projects                  |
| `/api/v1/projects/:id`        | GET    | Get project by id                 |
| `/api/v1/projects?name=<name>`| GET    | Find project by name              |
| `/api/v1/projects`            | POST   | Create project                    |
| `/api/v1/projects/:id`        | PUT    | Update project                    |
| `/api/v1/projects/:id`        | DELETE | Delete project                    |

## Troubleshooting

### Remove MongoDB data volume

List current volumes:

```
$ docker volume ls
```

Remove data volume:

```
$ docker volume rm crypto-wallet-api_mongo-data
```

Or

```
$ docker-compose down --volumes
```
