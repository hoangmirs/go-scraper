production:
	APP_RUN_MODE=${APP_RUN_MODE} bee run

env-setup:
	docker-compose -f docker-compose.dev.yml up -d

dev:
	make env-setup
	forego start
