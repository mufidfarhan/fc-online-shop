# Online Shop Project
1. Jalankan docker untuk PostgreSQL

```
docker run --name postgresql -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=123 -e POSTGRES_DB=online-shop -d -p 5434:5432 postgres:16
```

2. Export environtment variable yang dibutuhkan

```
export DB_URI=postgres://postgres:123@localhost:5434/online-shop?sslmode=disable
export ADMIN_SECRET=secret
```

3. Jalankan program

```
go run main.go
```

