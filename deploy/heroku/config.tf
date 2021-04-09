resource "heroku_config" "common" {
  vars = {
    APP_RUN_MODE = "prod"
  }
}

resource "heroku_app_config_association" "default" {
  app_id = heroku_app.default.id

  vars = heroku_config.common.vars
}
