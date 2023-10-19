## Golang migrate

- https://github.com/golang-migrate/migrate
- https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### Install

```bash
brew install golang-migrate
```

### Create migration

```bash
migrate create -ext sql -dir sql/migrations -seq init
```

- `-ext sql` - extension of the migration files
- `-dir sql/migrations` - directory where the migration files will be created
- `-seq init` - name of the migration file

### Run migrations

```bash
# apply:
migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose up

# rollback:
migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/goexpert" -verbose down
```

- 'up' - run all migrations
- 'down' - rollback all migrations