# Golang Restful API Boilerplate (go-chi)

Boilerplate api built using `go-chi`.

`Go Version: 1.22.2`

Table of contents
=================

* [Installation](#installation)
  * [Packagse](#packages)
* [Application](#application)
  * [Run Application](#run-api)
  * [Run Test](#run-test)
  * [Folder Structure](#folder-structure)

# Installation

### Packages
- [go-chi](https://github.com/go-chi/chi) - lighweight router for building Go HTTP services.
- [bun](https://github.com/uptrace/bun) - Lightweight Golang ORM (using PostgreSQL).
- [jwt](https://github.com/golang-jwt/jwt) - Go implementation of JSON Web Token Authentication.
- [yaml](https://gopkg.in/yaml.v3) - enables to decode YAML files to used for configuration.
- [gomail](https://gopkg.in/gomail.v2) - Golang SMTP mailer.
- [websocket](https://github.com/gorilla/websocket) - used for notifications and chat <span style="color: red">(*in progress*)</span>.

# Application

### Run API
```sh
go run .
```

### Run Test
```sh
go test ./tests
```

### Folder Structure
```
│   main.go                     #
├───conf                        # Configuration:
│       config.template.yml     #   *remove .template to be able to use configuration*
|                               #
├───controllers                 # Controllers:
│       AppControllers.go       #
|                               #
├───middlewares                 # Middlewares:
│       AppMiddlewares.go       #
|                               #
├───models                      # Models:
│       AppModels.go            #
|                               #
├───routers                     # Routers:
│       router.go               #
|                               #
├───services                    # Services:
|       AppServices.go          #
|                               #
├───sql                         # SQL:
|                               #
├───template                    # Templates:
│   │   404.html                #   this folder can be used to deploy compiled web application to avoid getting cors error
│   │   index.html              #   the 404 page is for when directing to api url using browser
│   ├───css                     #
│   └───emails                  #
│           welcome.html        #
├───tests                       # Tests:
|                               #
└───utils                       # Utilities:
        utility.go              #
```

##
[![GoDoc Widget]][GoDoc]
[![GitHub Widget]][My GitHub]

[GoDoc]: https://pkg.go.dev/github.com/go-chi/chi/v5
[GoDoc Widget]: https://img.shields.io/badge/references-go?style=flat&logo=go&labelColor=%235C5C5C&color=%23007D9C
[My GitHub]: https://github.com/XaiPhyr
[GitHub Widget]: https://img.shields.io/badge/XaiPhyr-github?style=flat&logo=github&labelColor=%235C5C5C&color=%235C5C5C
