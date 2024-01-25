<div align="center">
  <h1>Order Restaurant in go</h1>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/janapc/order-restaurant"/>
  <img alt="Language top" src="https://img.shields.io/github/languages/top/janapc/order-restaurant"/>
  <img alt="Repo size" src="https://img.shields.io/github/repo-size/janapc/order-restaurant"/>

<a href="#description">Description</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#usage">Usage</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#resources">Resources</a>

</div>

## Description

System to manager a restaurant and orders using rabbitmq and databases.

## Requirement

To this project your need:

- golang v1.21 [Golang](https://go.dev/)
- docker [Docker](https://www.docker.com/)

Inside **.docker** create a first file mysql_password.txt and put your password mysql and create a second file mysql_root_password and put your password root mysql.This configuration wil be used by docker-compose.yaml.

In root folder create a file **.env** with:

```env
MONGO_URL= //mongodb url
MONGO_DATABASE = //mongodb database name
MYSQL_URL= //mysql url
COLLECTION_MONGO_ORDERS= //mongodb collection
RABBITMQ_CONNECTION=//rabbitmq url
RABBITMQ_QUEUE_NAME=//queue name url
```

## Usage

Start Docker in your machine and run this commands in your terminal:

```sh
## up services
❯ docker compose up -d

## run this command to install dependencies:
❯ go mod tidy

## run this command to start api(localhost:3000/):
❯ go run cmd/main.go

```

Examples stay inside **test** folder

## Resources

- Golang
- Mysql
- Docker
- Mongo
- Fiber
- Rabbitmq

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
