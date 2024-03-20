#!/bin/bash

# Running docker-compose to build and start the film-api container
docker-compose up --build film-api

# Checking if the film-api container was successfully started
if [ $? -eq 0 ]; then
    echo "Film-api build completed successfully. Running migrations."
else
    echo "Error during film-api build. Exiting script."
    exit 1
fi

# Running the migrate command to apply migrations to the database
docker exec -i film-lib-db-1 psql -U postgres -d postgres < ./schema/000001_init.up.sql

# Checking if the migrate command was executed successfully
if [ $? -eq 0 ]; then
    echo "Migrations applied successfully to the database."
else
    echo "Error while running migrations."
fi