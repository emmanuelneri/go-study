# GO complete APP


#### Environment 

1. Start Postgres database 
```
docker run -it \
    --name postgres \
    -p 5432:5432 \
    -e POSTGRES_DB=goapp \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres \
    postgres:12-alpine
```    
2. Create tables 

- Access postgres `docker exec -it postgres psql -U postgres -d goapp`
    - Create order table: ``create table if not exists sales_order (id integer primary key generated always as identity, customer VARCHAR(200), total numeric(19, 2));``

#### RUN

RUN project ``go run src/main.go``