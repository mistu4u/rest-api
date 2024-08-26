FROM postgres:latest

# Copy SQL script into the Docker container
COPY init.sql /docker-entrypoint-initdb.d/

# Expose PostgreSQL port
EXPOSE 5432