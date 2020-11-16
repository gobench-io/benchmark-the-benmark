read CLIENT
echo monitor benchmark client $CLIENT
WORKER=nginx
BENCHMARK_CONTAINER=benchmark-client

docker ps -a --filter name=$WORKER -aq | xargs -r docker rm -f
docker ps -a --filter name=$BENCHMARK_CONTAINER -aq | xargs -r docker rm -f

cd $CLIENT && docker-compose --compatibility up -d --b


echo streaming resource from benchmark client  $CLIENT  ....
docker stats $BENCHMARK_CONTAINER > $CLIENT.top 