# Itrack
Itrack is a web system for tracking IN's and OUT's of authorized people to a physical infrastructure.

# Backend 
This repository contains the backend side written in Golang, it can be use as source code only and compiled manually, also it can be built in a container (Dockerfile provided). Additionally there is a docker-compose yml file for fast dev deployment with a mongo service included.

### Installation

Backend requires [Golang](https://golang.org/) v1.1+ to run.

Install the dependencies and devDependencies and start the server.

Download dependencies using go mod
```sh
$ go mod download
```

For production, build the application and run the executable

```sh
$ go build -o main .
$ ./main
```

For development, run command
```sh
$ go run main.go db.go rand.go
```