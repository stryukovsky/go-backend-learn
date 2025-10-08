docker run -d --name postgresql_database -e POSTGRES_USER=user -e POSTGRES_PASSWORD=pass -e POSTGRES_DB=db -p 5432:5432 postgres 2>/dev/null
sleep 3
docker restart postgresql_database 2>/dev/null

