## db migration

```bash
$ migrate create -ext sql -seq -dir db/migrations user_profiles
$ migrate create -ext sql -seq -dir db/migrations user_credentials
```

### migrate up

```bash
$ migrate -path ./db/migrations -database "postgresql://my_user:pass@localhost:5432/user-service?sslmode=disable" up 2
```

### migrate down

```bash
$ migrate -path ./db/migrations -database "postgresql://my_user:pass@localhost:5432/user-service?sslmode=disable" down 2
```

## proto

```bash
$ protoc --proto_path=. --micro_out=. --go_out=. .\src\proto\*.proto
```

## run

```bash
$ go run main.go allSvc
```