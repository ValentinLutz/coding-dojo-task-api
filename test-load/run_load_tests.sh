#/bin/bash
set -e

run_load_test() {
  app_name=$1
  use_memory=$2

  if $use_memory;
  then
    test_id="$app_name-memory"
  else
    test_id="$app_name-postgres"
  fi

  printf "Run load test for $test_id\n"

  export APP_NAME=$app_name
  export USE_MEMORY=$use_memory
  export TEST_ID=$test_id

  # start the app
  docker compose -f docker-compose.app.yaml up -d \
    --force-recreate \
    --wait
  # wait for the app to start
  sleep 20

  docker run -it \
    --rm \
    --name k6-$test_id \
    --volume ./script.js:/k6/script.js \
    --network host \
    --env K6_PROMETHEUS_RW_TREND_AS_NATIVE_HISTOGRAM=true \
    grafana/k6:0.45.0 \
    run \
    --out experimental-prometheus-rw \
    --tag testid=$test_id \
    /k6/script.js \
    || true

  # stop the app
  docker compose -f docker-compose.app.yaml down
}

docker compose -f docker-compose.test.yaml up -d \
  --wait

for use_memory in \
  true \
  false

do
  for app_name in \
    "task-app-go-chi" \
    "task-app-go-fiber" \
    "task-app-java-javalin" \
    "task-app-java-quarkus-reactive" \
    "task-app-java-quarkus-resteasy" \
    "task-app-java-spring-web-mvc"
    "task-app-java-spring-webflux" \
    "task-app-rust-axum" \
  do
    run_load_test $app_name $use_memory
  done
done
