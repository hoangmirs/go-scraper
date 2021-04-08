# Go Scraper

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

Step 1: Copy the variable file and update the variables

```
cp terraform.tfvars.sample terraform.tfvars
```

Step 2: Initialize Terraform

```
terraform init
```

Step 3: Generate an execution plan

```
terraform plan
```

Step 5: Execute the generated plan

```
terraform apply
```

Step 6: Build the application and push to heroku
You can check `.github/workflows/deploy.yml` workflow for more details

## About

This project is created to complete **Web Certification Path** using **Go** at [Nimble](https://nimblehq.co)
