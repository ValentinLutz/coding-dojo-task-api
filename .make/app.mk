app/config/config.yaml: ./config/config.yaml ## Copy none-dev config to app directory
	install -D config/config.yaml app/config/config.yaml

app/serve/openapi/tasks.yaml: api-definition/tasks.yaml ## Copy tasks open api definition to app
	install -D api-definition/tasks.yaml app/serve/openapi/tasks.yaml

app/api/task/tasks.gen.go: api-definition/tasks.yaml ## Generate tasks server from open api definition
	oapi-codegen -generate types \
		-package task \
		./api-definition/tasks.yaml  > app/api/task/tasks.gen.go

app.run:: app/config/config.yaml app/serve/openapi/tasks.yaml app/api/task/tasks.gen.go ## Run the app
	cd app && \
		go run main.go

app.build:: app/serve/openapi/tasks.yaml app/api/task/tasks.gen.go ## Build the app into an executable
	cd app && \
		go build

app.lint:: app/api/task/tasks.gen.go ## Runs linters against go code
	cd app && \
		golangci-lint run
