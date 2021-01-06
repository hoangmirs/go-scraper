.PHONY: env-setup dev test

install-dependencies:
	go get github.com/beego/bee/v2
	go get github.com/ddollar/forego
	go mod tidy

env-setup:
	docker-compose -f docker-compose.dev.yml up -d
	npm install

dev:
	make env-setup
	forego start

test:
	docker-compose -f docker-compose.test.yml up -d
	APP_RUN_MODE=test go test -v -p 1 -count=1 ./...
	docker-compose -f docker-compose.test.yml down
