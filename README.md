# bmstu-news
News parsing service from the official BMSTU website

## About
A parsing service for requesting news from the MSTU website and writing them to the database 

## Commands
#### Run
Runs application
```shell
make run
```
First actions to run application
```shell
make local
docker-compose up -d
make migrate
make run
```
#### First
Write all news into database
```shell
make first
```
#### Migrate
Create migrations
```shell
make migrate
```
#### Local
Creates environment file
```shell
make local
```


