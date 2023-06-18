#/bin/bash
set -e

docker compose -f docker-compose.yaml up -d \
  --wait

docker build \
  -f xk6-influxdb/Dockerfile \
  -t xk6-influxdb:local .

for use_memory in "true" "false"
do
  for app_name in "task-app-java-quarkus-reactive" "task-app-java-spring-web-mvc" "task-app-go-chi" "task-app-go-fiber"
  do
    printf "Run load test for app '$app_name' with memory '$use_memory'\n"
    ./run_load_test.sh $app_name $use_memory
  done
done