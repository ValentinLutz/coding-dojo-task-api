.DEFAULT_GOAL := help

help::
	@egrep -h '\s##\s' $(MAKEFILE_LIST) \
		| awk -F':.*?## | \\| ' '{printf "\033[36m%-38s \033[37m %-68s \033[35m%s \n", $$1, $$2, $$3}'


export PROFILE ?= none-local
export FLYWAY_USER ?= test
export FLYWAY_PASSWORD ?= test

app.gen:: ## Run go generate
	cd app && \
		go generate ./...

app.run:: app.gen## Run the app
	cd app && \
		go run main.go

app.build:: app.gen ## Build the app into an executable
	cd app && \
		go build


test.unit:: app.gen## Run the unit tests
	cd app && \
		go test -race -cover ./...

test.load:: ## Run load tests
	k6 run test-load/script.js 


database.migrate:: ## Migrate database | PROFILE, FLYWAY_USER, FLYWAY_PASSWORD
	cd migration-database && \
		flyway clean \
		migrate \
		-configFiles=${PROFILE}.properties \
		-user=${FLYWAY_USER} \
		-password=${FLYWAY_PASSWORD}


docker.up:: ## Start containers 
	docker compose -f deployment-docker/docker-compose.yaml \
		up --force-recreate --build