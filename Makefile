build_prod:
	-docker rm backend-jobstreet-application-service-prod
	-docker rmi backend-jobstreet-application-service-prod
	docker-compose -f docker-compose.prod.yml build

# Shouldn't run. big image size ~ 1GB
build_dev:
	-docker rm backend-jobstreet-application-service-dev
	-docker rmi backend-jobstreet-application-service-dev
	docker-compose -f docker-compose.dev.yml build

run_prod:
	docker run -p 4000:4000 --name backend-jobstreet-application-service-prod backend-jobstreet-application-service-prod

run_dev:
	docker run -p 4000:4000 --name backend-jobstreet-application-service-dev backend-jobstreet-application-service-dev

start_prod:
	docker start backend-jobstreet-application-service-prod

start_dev:
	docker start backend-jobstreet-application-service-dev

build_run_prod:
	make build_prod
	make run_prod

.PHONY: build_prod build_run_dev run_prod start_prod start_dev build_run_prod