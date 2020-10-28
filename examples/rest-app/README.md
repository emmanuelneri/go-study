# Rest APP


#### Environment 

Start Postgres database 
```
docker run -it \
    --name postgres \
    -p 5432:5432 \
    -e POSTGRES_DB=goapp \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres \
    postgres:12-alpine
```    

#### RUN

RUN project ``go run cmd/main.go``


#### Run profile

- Heap `go tool pprof http://localhost:8080/debug/pprof/heap`
- build project `go tool pprof http://localhost:8080/debug/pprof/profile?seconds=120`