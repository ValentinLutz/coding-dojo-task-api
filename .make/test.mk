test.unit::  app/serve/openapi/tasks.yaml app/api/task/tasks.gen.go ## Run the unit tests
	cd app && \
		go test ./...

test:: test.unit ## Run all tests