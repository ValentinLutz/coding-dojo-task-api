docker.build:: app/serve/openapi/tasks.yaml app/api/task/tasks.gen.go ## Build container images | DOCKER_REGISTRY, DOCKER_REPOSITORY, PROJECT_NAME, VERSION
	docker build \
    		-t ${DOCKER_REGISTRY}/${DOCKER_REPOSITORY}/${PROJECT_NAME}:${VERSION} \
    		-t ${DOCKER_REGISTRY}/${DOCKER_REPOSITORY}/${PROJECT_NAME}:latest \
    		app

docker.up:: docker.build ## Start containers | PROJECT_NAME
	docker-compose -p ${PROJECT_NAME} \
		-f deployment-docker/docker-compose.yaml \
		up \
		--force-recreate

docker.down:: ## Shutdown containers | PROJECT_NAME
	docker-compose -p ${PROJECT_NAME} \
		-f deployment-docker/docker-compose.yaml \
		down