.DEFAULT_GOAL := help

help::
	@egrep -h '\s##\s' $(MAKEFILE_LIST) \
		| awk -F':.*?## | \\| ' '{printf "\033[36m%-38s \033[37m %-68s \033[35m%s \n", $$1, $$2, $$3}'

build.images:: ## Build docker images
	make -C go-chi docker.image && \
		make -C go-fiber docker.image && \
		make -C java-quarkus-resteasy docker.image && \
		make -C java-quarkus-reactive docker.image && \
		make -C java-spring-web-mvc docker.image && \
		make -C java-spring-webflux docker.image && \
		make -C java-javalin docker.image && \
		make -C rust-axum docker.image

test.load:: ## Run load tests
	cd test-load && \
		bash run_load_tests.sh

test.functional:: ## Run functional tests
	cd test-functional && \
		go test -count=1 ./...

start.container:: ## Start docker containers
	docker compose -f test-load/docker-compose.test.yaml up -d

stop.container:: ## Stop docker containers
	docker compose -f test-load/docker-compose.app.yaml down && \
		docker compose -f test-load/docker-compose.test.yaml down
