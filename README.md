<div align="center">
  <h1>Movie api grpc</h1>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/janapc/grpc-movies-go"/>
  <img alt="Language top" src="https://img.shields.io/github/languages/top/janapc/grpc-movies-go"/>
  <img alt="Repo size" src="https://img.shields.io/github/repo-size/janapc/grpc-movies-go"/>

<a href="#project">Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#requirement">Requirement</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#run-project">Run Project</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
<a href="#technologies">Technologies</a>

</div>

## Project

Api to manager movies.

## Requirement

To this project your need:

- golang v1.21 [Golang](https://go.dev/)
- docker [Docker](https://www.docker.com/)

In root folder create a file **.env** with:

```env
MONGO_URI= //database connection
```

## Run Project

Start Docker in your machine and run this commands in your terminal:

```sh
## up mongodb
❯ docker compose up -d

## run this command to install dependencies:
❯ go mod tidy

## run this command to start api(localhost:50051):
❯ go run cmd/grpc-server/main.go

```

## Technologies

- golang
- grpc
- docker
- mongoDB

<div align="center">

Made by Janapc 🤘 [Get in touch!](https://www.linkedin.com/in/janaina-pedrina/)

</div>
