.PHONY: env-setup test

install-dependencies:
	go get github.com/beego/bee

env-setup:
	docker-compose -f docker-compose.dev.yml up -d
	npm install

dev:
	make env-setup
	forego start

test:
	go test -v -p 1 ./...

production:
	APP_RUN_MODE=${APP_RUN_MODE} bee run
