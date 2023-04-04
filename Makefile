.DEFAULT_GOAL := help

help::
	@egrep -h '\s##\s' $(MAKEFILE_LIST) \
		| awk -F':.*?## | \\| ' '{printf "\033[36m%-38s \033[37m %-68s \033[35m%s \n", $$1, $$2, $$3}'


export PROFILE ?= none-local
export FLYWAY_USER ?= test
export FLYWAY_PASSWORD ?= test


dep.oapi-codegen:: # Install oapi-codegen with go install
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4


app/incoming/taskapi/server.gen.go: dep.oapi-codegen api-definition/task_api.yaml api-definition/server.app.yaml ## Generate task api server from open api definition
	oapi-codegen --config api-definition/server.app.yaml \
		api-definition/task_api.yaml

app/incoming/taskapi/types.gen.go: dep.oapi-codegen api-definition/task_api.yaml api-definition/types.app.yaml ## Generate task api types from open api definition
	oapi-codegen --config api-definition/types.app.yaml \
		api-definition/task_api.yaml


app.run:: app/incoming/taskapi/server.gen.go app/incoming/taskapi/types.gen.go  ## Run the app
	cd app && \
		go run main.go

app.build:: app/incoming/taskapi/server.gen.go app/incoming/taskapi/types.gen.go ## Build the app into an executable
	cd app && \
		go build


test.unit::  app/incoming/taskapi/server.gen.go app/incoming/taskapi/types.gen.go ## Run the unit tests
	cd app && \
		go test -race -cover ./...

test.load:: ## Run load tests
	docker run -it \
		--rm \
		--volume ${PWD}/test-load:/k6 \
		--network host  \
        grafana/k6:0.39.0 \
		run /k6/script.js \


docker.up:: app/incoming/taskapi/server.gen.go app/incoming/taskapi/types.gen.go ## Start containers 
	docker compose -f deployment-docker/docker-compose.yaml \
		up --force-recreate --build