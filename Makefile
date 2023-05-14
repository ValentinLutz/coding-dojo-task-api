.DEFAULT_GOAL := help

help::
	@egrep -h '\s##\s' $(MAKEFILE_LIST) \
		| awk -F':.*?## | \\| ' '{printf "\033[36m%-38s \033[37m %-68s \033[35m%s \n", $$1, $$2, $$3}'

export GOPROXY=direct

export PROFILE ?= none-local
export FLYWAY_USER ?= test
export FLYWAY_PASSWORD ?= test


dep.oapi-codegen:: # Install oapi-codegen with go install
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4


app-chi/incoming/taskapi/server.gen.go: dep.oapi-codegen api-definition/task_api.yaml api-definition/server.app-chi.yaml ## Generate task api server from open api definition
	oapi-codegen --config api-definition/server.app-chi.yaml \
		api-definition/task_api.yaml

app-chi/incoming/taskapi/types.gen.go: dep.oapi-codegen api-definition/task_api.yaml api-definition/types.app-chi.yaml ## Generate task api types from open api definition
	oapi-codegen --config api-definition/types.app-chi.yaml \
		api-definition/task_api.yaml


app.chi.run:: app-chi/incoming/taskapi/server.gen.go app-chi/incoming/taskapi/types.gen.go  ## Run the app
	cd app-chi && \
		go run -race main.go

docker.chi.up:: app-chi/incoming/taskapi/server.gen.go app-chi/incoming/taskapi/types.gen.go ## Start containers 
	docker compose -f deployment-docker/docker-compose.chi.yaml \
		up --force-recreate --build

app-fiber/incoming/taskapi/types.gen.go: dep.oapi-codegen api-definition/task_api.yaml api-definition/types.app-fiber.yaml ## Generate task api types from open api definition
	oapi-codegen --config api-definition/types.app-fiber.yaml \
		api-definition/task_api.yaml

app.fiber.run:: app-fiber/incoming/taskapi/types.gen.go  ## Run the app
	cd app-fiber && \
		go run -race main.go

docker.fiber.up:: app-fiber/incoming/taskapi/types.gen.go ## Start containers 
	docker compose -f deployment-docker/docker-compose.fiber.yaml \
		up --force-recreate --build

test.load:: ## Run load tests
	docker run -it \
		--rm \
		--volume ${PWD}/test-load:/k6 \
		--network host  \
        grafana/k6:0.39.0 \
		run /k6/script.js \

