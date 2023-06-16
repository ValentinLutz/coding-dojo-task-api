.DEFAULT_GOAL := help

help::
	@egrep -h '\s##\s' $(MAKEFILE_LIST) \
		| awk -F':.*?## | \\| ' '{printf "\033[36m%-38s \033[37m %-68s \033[35m%s \n", $$1, $$2, $$3}'

export GOPROXY=direct

dep.oapi-codegen:: # Install oapi-codegen with go install
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4


app/incoming/taskapi/types.gen.go: dep.oapi-codegen api-definition/task_api.yaml api-definition/types.app.yaml ## Generate task api types from open api definition
	oapi-codegen --config api-definition/types.app.yaml \
		api-definition/task_api.yaml

app.run:: app/incoming/taskapi/types.gen.go  ## Run the app
	cd app && \
		go run main.go

docker.up:: app/incoming/taskapi/types.gen.go ## Start containers 
	docker compose -f deployment-docker/docker-compose.yaml \
		up --force-recreate --build
