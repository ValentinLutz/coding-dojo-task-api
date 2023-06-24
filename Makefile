.DEFAULT_GOAL := help

help::
	@egrep -h '\s##\s' $(MAKEFILE_LIST) \
		| awk -F':.*?## | \\| ' '{printf "\033[36m%-38s \033[37m %-68s \033[35m%s \n", $$1, $$2, $$3}'

build.images:: ## Build docker images
	docker compose -f go-chi/deployment-docker/docker-compose.yaml build && \
		docker compose -f go-fiber/deployment-docker/docker-compose.yaml build && \
		docker compose -f java-quarkus-resteasy/deployment-docker/docker-compose.yaml build && \
		docker compose -f java-quarkus-reactive/deployment-docker/docker-compose.yaml build && \
		docker compose -f java-spring-web-mvc/deployment-docker/docker-compose.yaml build && \
		docker compose -f java-spring-webflux/deployment-docker/docker-compose.yaml build && \
		docker compose -f java-javalin/deployment-docker/docker-compose.yaml build

test.load:: ## Run load tests
	cd test-load && \
		bash run_load_tests.sh

test.functional:: ## Run functional tests
	cd test-functional && \
		go test -count=1 ./...

stop.container:: ## Stop docker containers
	docker compose -f test-load/docker-compose.app.yaml down && \
		docker compose -f test-load/docker-compose.test.yaml down
