# podman pull quay.io/fedora/redis-7
podman start redis_database || podman run -d --name redis_database -p 6379:6379 -e REDIS_PASSWORD=redis redis 
