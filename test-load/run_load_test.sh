app_name=$1
use_memory=$2

if [ $use_memory ]; then
    script_name=script-memory.js
else
    script_name=script-postgres.js
fi

cp script.js $script_name

export APP_NAME=$app_name
export USE_MEMORY=$use_memory

# start the app
docker compose up -d --force-recreate --wait
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

docker run -it \
   --rm \
   --volume ${PWD}:/k6 \
		--network host \
		grafana/k6:0.44.1 \
		run /k6/$script_name

# stop the app
docker compose down \
  --remove-orphans
