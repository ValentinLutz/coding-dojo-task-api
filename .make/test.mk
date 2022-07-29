test.unit::  app/serve/openapi/tasks.yaml app/api/task/tasks.gen.go ## Run the unit tests
	cd app && \
		go test ./...

test.load:: ## Run load tests
	docker run -it \
		--rm \
		--volume ${PWD}/test-load:/k6 \
		--network coding-dojo-api-golang \
        grafana/k6:0.39.0 \
		run /k6/script.js \

test:: test.unit ## Run all tests