<div align="center">
  <h1>Order Restaurant in go</h1>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/janapc/order-restaurant"/>
  <img alt="Language top" src="https://img.shields.io/github/languages/top/janapc/order-restaurant"/>
  <img alt="Repo size" src="https://img.shields.io/github/repo-size/janapc/order-restaurant"/>

<a href="#project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#run-project">Run Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#request-api">Request API</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#technologies">Technologies</a>

</div>

## Project

Api to manager order in restaurant using clean arch.

## Requirement

To this project your need:

- golang v1.21 [Golang](https://go.dev/)
- docker [Docker](https://www.docker.com/)

Inside **.docker** create a first file mysql_password.txt and put your password mysql and create a second file mysql_root_password and put your password root mysql.This configuration wil be used by docker-compose.yaml.

In root folder create a file **.env** with:

```env
MYSQL_URL= # mysqldb url
MONGO_URL= # mongodb url
```

## Run Project

Start Docker in your machine and run this commands in your terminal:

```sh
## up mysql
❯ docker compose up -d

## run this command to install dependencies:
❯ go mod tidy

## run this command to start api(localhost:3000/):
❯ go run cmd/main.go

```

## Request API

Examples stay inside **test** folder

## Technologies

- golang
- mysql
- docker
- mongo

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
