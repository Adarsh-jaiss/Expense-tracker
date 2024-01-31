# Expense tracker backend

### packages used

- Go fiber : `go get github.com/lib/pq`
- Psql : `go get github.com/lib/pq`

## docker + Postgres

#### Setting up postgress locally
- Postgres docker image : `docker pull postgres`
- Run a docker container : `docker run --name my-postgres-container -e POSTGRES_PASSWORD=<your-postgres-password> -p 5432:5432 -d postgres`
- Access the Postgress database using psql client : `docker run -it --rm --name psql-client postgres psql -h <your-host-ip> -U postgres -d postgres`

Note : if you don't know the exact IP address of your Postgres container, use this commmand to find it : `docker inspect my-postgres-container`
