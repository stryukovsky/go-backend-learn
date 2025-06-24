podman pull quay.io/fedora/redis-7
podman run -d --name redis_database -e REDIS_PASSWORD=redis fedora/redis-7 -p 6379:6379
