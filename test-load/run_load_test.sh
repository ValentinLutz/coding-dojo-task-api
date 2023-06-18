#/bin/bash
set -e

app_name=$1
use_memory=$2

if [ $use_memory ];
then
  test_id="$app_name-memory"
else
  test_id="$app_name-postgres"
fi

export APP_NAME=$app_name
export USE_MEMORY=$use_memory

# start the app
docker compose -f docker-compose.test.yaml up -d \
  --force-recreate \
  --wait
# wait for the app to start
sleep 30

# run the load test
# docker run -it \
#     --rm \
#     --volume ${PWD}:/k6 \
# 		--network host \
# 		--env K6_CLOUD_TOKEN=XXX \
# 		grafana/k6:0.44.1 \
# 		run --out=cloud /k6/$script_name

#docker run -it \
#   --rm \
#   --volume ${PWD}:/k6 \
#		--network host \
#		grafana/k6:0.44.1 \
#		run /k6/$script_name

docker run -it \
  --rm \
  --volume ${PWD}:/k6 \
  --network host \
  --env K6_INFLUXDB_ORGANIZATION=monke \
  --env K6_INFLUXDB_BUCKET=coding_dojo \
  --env K6_INFLUXDB_TOKEN=test \
  --env K6_INFLUXDB_ADDR=http://localhost:8086 \
  --env K6_INFLUXDB_CONCURRENT_WRITES=20 \
  --env K6_INFLUXDB_PUSH_INTERVAL=10s \
  xk6-influxdb:local \
  run \
  --tag testid=$test_id \
  -o xk6-influxdb \
  /k6/script.js

# stop the app
docker compose -f docker-compose.test.yaml down
