docker rm -f postgresql_database redis_database
sleep 1
docker volume prune -a 
sleep 1
./startdb.sh
sleep 10
go run . migrate
sleep 1
go run . load
sleep 1
./startredis.sh
