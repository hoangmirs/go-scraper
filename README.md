# Go Scraper

## Prerequisite

- [Go - 1.15](https://golang.org/doc/go1.15)
- Docker: check for the [installation guide](https://www.docker.com/products/docker-desktop) for your platform

## Usage

Clone the repository

`git clone git@github.com:nimblehq/git-template.git`

#### Install development dependencies

This project uses the following dependencies:

- [Bee - Bee CLI](https://github.com/beego/bee)
- [Forego - Foreman in Go](https://github.com/ddollar/forego)

Install them by running:

```sh
  make install-dependencies
```

#### Run the Go application for development

```sh
  make dev
```

The application will be run on http://localhost:8080

#### Run tests

```sh
  make env-setup
  make test
```

## About

This project is created to complete **Web Certification Path** using **Go** at [Nimble](https://nimblehq.co)
