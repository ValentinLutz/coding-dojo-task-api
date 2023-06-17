.DEFAULT_GOAL := help

help::
	@egrep -h '\s##\s' $(MAKEFILE_LIST) \
		| awk -F':.*?## | \\| ' '{printf "\033[36m%-38s \033[37m %-68s \033[35m%s \n", $$1, $$2, $$3}'

# test.load:: ## Run load tests
#	docker run -it \
#		--rm \
#		--volume ${PWD}/test-load:/k6 \
#		--network host  \
#       grafana/k6:0.44.1 \
#		run /k6/script.js \

test.load:: ## Run load tests
	docker run -it \
		--rm \
		--volume ${PWD}/test-load:/k6 \
		--network host  \
        ghcr.io/szkiba/xk6-dashboard:0.4.3 \
		run --out json=/k6/test_result.json /k6/script.js

test.report:: ## Run load tests
	docker run -it \
		--rm \
		--volume ${PWD}/test-load:/k6 \
		--network host  \
        ghcr.io/szkiba/xk6-dashboard:0.4.3 \
		dashboard replay /k6/test_result.json