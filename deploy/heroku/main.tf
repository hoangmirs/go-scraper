terraform {
  required_providers {
    heroku = {
      source  = "heroku/heroku"
      version = "~> 4.0"
    }
  }
}

provider "heroku" {
  email   = var.heroku_email
  api_key = var.heroku_api_key
}


resource "heroku_app" "default" {
  name   = var.app_name
  region = var.app_region
  stack  = var.app_stack
}

resource "heroku_addon" "postgresql" {
  app  = heroku_app.default.name
  plan = "heroku-postgresql:${var.postgresql_plan}"
}

resource "heroku_addon" "redis" {
  app  = heroku_app.default.name
  plan = "heroku-redis:${var.redis_plan}"
}
