#/bin/bash

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

# start the app
docker compose -f docker-compose.app.yaml up -d \
  --force-recreate \
  --wait
# wait for the app to start
sleep 30

docker run -it \
  --rm \
  --volume ./script.js:/k6/script.js \
  --network host \
  --env K6_PROMETHEUS_RW_TREND_AS_NATIVE_HISTOGRAM=true \
  grafana/k6:0.45.0 \
  run \
  --out experimental-prometheus-rw \
  --tag testid=$test_id \
  /k6/script.js

# stop the app
docker compose -f docker-compose.app.yaml down
