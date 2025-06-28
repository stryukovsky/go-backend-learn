podman pull quay.io/sclorg/postgresql-15-c9s
podman start postgresql_database || podman run -d --name postgresql_database -e POSTGRESQL_USER=user -e POSTGRESQL_PASSWORD=pass -e POSTGRESQL_DATABASE=db -p 5432:5432 postgresql-15-c9s
