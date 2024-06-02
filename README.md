# backend-user-management-service
This repo is Candidate Service of jobstreet application backend.

## Deploy

1. ```docker build -t thanhquy1105/backend-jobstreet-candidate-service-prod:latest .```
2. ```docker push thanhquy1105/backend-jobstreet-candidate-service-prod```
3. ```docker pull thanhquy1105/backend-jobstreet-candidate-service-prod:latest```
4. ```docker run --name backend-jobstreet-candidate-service-prod --network jobstreet-network -p 4002:4002 -e DB_SOURCE="postgresql://admin:admin@postgres:5432/candidate_service_jobstreet?sslmode=disable" -d thanhquy1105/backend-jobstreet-candidate-service-prod:latest```

## RUN

1. Run eventstoreDB (if it's the first run)
    ```make esdb```
or (from the second run)
    ```make start_esdb```
2. Run PostgresDB (if it's the first run)
    ```make run_postgres```
    ```make createdb```
or (from the second run)
    ```make start_postgres```
3. Run server
    ```go run main.go```

