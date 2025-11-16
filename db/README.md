# Database 

Using PSQL 18, deployed via Docker

Not required for the ETL binary, but for any usecases that goes beyond the csv file, database is good.

## Configuration

There are 5 environment variables required for Postgres

`SJSUBARK_PSQL_PASSWORD`

`SJSUBARK_PSQL_USER`

`SJSUBARK_PSQL_DB`

`SJSUBARK_PSQL_PORT`

`SJSUBARK_PSQL_HOST`

## RESTAPI

I've created a lightweight server that serves request to the SQL server. It can be found in server.go and /datastore.

The purpose is so the frontend app isn't directly accessing the data layer.
