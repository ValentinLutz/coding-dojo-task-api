#/bin/bash
set -e

docker compose -f docker-compose.test.yaml up -d \
  --wait

for use_memory in \
  true \
  false

do
  for app_name in \
  "task-app-java-quarkus-resteasy" \
  "task-app-java-spring-webflux" \
  "task-app-java-javalin" \
  "task-app-java-quarkus-reactive" \
  "task-app-java-spring-web-mvc" \
  "task-app-go-chi" \
  "task-app-go-fiber"
  do
    ./run_load_test.sh $app_name $use_memory
  done
done