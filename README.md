# Go Scraper

## Demo

- [Staging](https://go-scraper-staging.herokuapp.com/)
- [Production](https://go-scraper-prod.herokuapp.com/)
- [Postman collection](https://documenter.getpostman.com/view/209740/TzJoELk2)

## Prerequisite

- [Go - 1.15](https://golang.org/doc/go1.15)
- [Node - 14.15.4](https://nodejs.org/en/)
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
  make test
```

#### Deploy to Heroku with Terraform

##### Prerequisites

- [Heroku CLI](https://devcenter.heroku.com/articles/heroku-cli) latest version
- [Terraform](https://www.terraform.io/downloads.html)

To deploy the application to Heroku with Terraform, we need to create the Heroku API Key first:

```bash
$ heroku authorizations:create --description <api key description>
```

And then, move to the `deploy/heroku` folder and run the following steps:

_Step 1:_ Copy the variable file and update the variables

```sh
$ cp terraform.tfvars.sample terraform.tfvars
```

_Step 2:_ Initialize Terraform

```sh
$ terraform init
```

_Step 3:_ Generate an execution plan

```sh
$ terraform plan
```

_Step 5:_ Execute the generated plan

```sh
$ terraform apply
```

_Step 6:_ Build the application and push to heroku

You can check `.github/workflows/deploy.yml` workflow for more details

_Make sure you set the following Github secrets before deploying the application:_

```
DOCKER_TOKEN       # Docker token
HEROKU_API_KEY     # Heroku OAuth token
HEROKU_APP_PROD    # Heroku app name for production
HEROKU_APP_STAGING # Heroku app name for staging
```

## About

This project is created to complete **Web Certification Path** using **Go** at [Nimble](https://nimblehq.co)
