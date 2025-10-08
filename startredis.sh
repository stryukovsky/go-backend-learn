docker run -d --name redis_database -p 6379:6379 -e REDIS_PASSWORD=redis redis 2>/dev/null
sleep 3
docker restart redis_database 2>/dev/null

