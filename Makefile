DB_URL=postgresql://admin:admin@localhost:5432/candidate_service_jobstreet?sslmode=disable

docker_compose_build:
	docker-compose build

docker_build:
	docker buildx build -t thanhquy1105/backend-jobstreet-candidate-service-prod:latest .

docker_push:
	docker push thanhquy1105/backend-jobstreet-candidate-service-prod

docker_build_run:
	docker-compose up

# generate a new migration
new_migrate:
	migrate create -ext sql -dir db/migration -seq $(name)

# run postgres container with network 
run_postgres:
	-docker network create jobstreet-network
	docker run --name postgres --network jobstreet-network -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:13.12 

start_postgres:
	docker start postgres

build_app:
	docker build -t thanhquy1105/backend-jobstreet-candidate-service-prod:latest .

run_app:
	docker run --name backend-jobstreet-candidate-service-prod --network jobstreet-network -p 4002:4002 -e DB_SOURCE="postgresql://admin:admin@postgres:5432/candidate_service_jobstreet?sslmode=disable" thanhquy1105/backend-jobstreet-candidate-service-prod:latest

start_app:
	docker start backend-jobstreet-candidate-service-prod

push_app:
	docker push thanhquy1105/backend-jobstreet-candidate-service-prod

# create candidate_service_jobstreet database on postgres container
createdb:
	docker exec -it postgres createdb --username=admin --owner=admin candidate_service_jobstreet

# drop candidate_service_jobstreet database on postgres container
dropdb:
	docker exec -it postgres dropdb --username=admin candidate_service_jobstreet

# migrate candidate_service_jobstreet database from app to postgres container
migrate:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

# generate queries to golang code
sqlc:
	docker run --rm -v "${CURDIR}:/src" -w /src sqlc/sqlc:1.20.0 generate

# generate swagger
swagger:
	swag init --parseDependency -g main.go

# run test
test:
	go test -v -cover -short ./...

esdb:
	docker run --name esdb-node -it -p 2113:2113 -p 1113:1113 \
    	eventstore/eventstore:21.6.0-buster-slim --insecure --run-projections=All \
    	--enable-external-tcp --enable-atom-pub-over-http

start_esdb:
	docker start esdb-node

.PHONY: build_run_prod new_migrate run_postgres migrate dropdb createdb start_postgres sqlc evans swagger proto esdb