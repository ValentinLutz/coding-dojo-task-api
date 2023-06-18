.DEFAULT_GOAL := help

help::
	@egrep -h '\s##\s' $(MAKEFILE_LIST) \
		| awk -F':.*?## | \\| ' '{printf "\033[36m%-38s \033[37m %-68s \033[35m%s \n", $$1, $$2, $$3}'


app.run:: ## Run the app
	cd app && \
		./mvnw compile spring-boot:run

docker.up:: ## Start containers
	docker compose -f deployment-docker/docker-compose.yaml \
		up --force-recreate --build
