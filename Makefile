.PHONY: env-setup dev test production

install-dependencies:
	go get github.com/beego/bee
	go get github.com/ddollar/forego
	go mod tidy

env-setup:
	docker-compose -f docker-compose.dev.yml up -d
	npm install

dev:
	make env-setup
	forego start

test:
	go test -v -p 1 ./...

production:
	bin/start
