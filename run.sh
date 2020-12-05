echo input benchmark client: 
read CLIENT
echo monitor benchmark client $CLIENT
WORKER=nginx
BENCHMARK_CONTAINER=benchmark-client
BENCHMARK_BASE=benchmark-base

# put current date as yyyy-mm-dd HH:MM:SS in $date
run=$(date '+%Y-%m-%d-%H:%M:%S')

docker ps -a --filter name=$WORKER -aq | xargs -r docker rm -f
docker ps -a --filter name=$BENCHMARK_CONTAINER -aq | xargs -r docker rm -f
docker ps -a --filter name=$BENCHMARK_BASE -aq | xargs -r docker rm -f

cd $CLIENT && docker-compose --compatibility up -d --b


echo streaming resource from benchmark client  $CLIENT  ....
docker stats  > $CLIENT_$run.top 
docker logs -f benchmark-client
