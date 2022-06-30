### Step 1

Install golang-migrate to start creating migration files

```bash
brew install golang-migrate
```

Create the first sql migration file

```bash
migrate create -ext sql -dir db/migrations -seq init_schema
```

Add your sql schemas into up.sql and down.sql files inside db/migrations folder

Fire up a postgres container

```bash
make postgres
```

Create a database

```bash
make createdb
```

Run the migration files

```bash
make migrateup
```

### Step 2 sqlc and generating go functions

```bash
brew install sqlc
```

```bash
sqlc init
```

Make sure you configure sqlc.yaml file

Create a query file that represents a query to a table in query folder e.g: query/authors.sql

```sql
-- name: CreateAuthor :execresult
INSERT INTO authors
    (
    name, bio
    )
VALUES
    (
        $1, $2
)
RETURNING *;
```

Now lets generate our sql-go files

```bash
make sqlc
```

### step 3 writing tests

```bash
go get github.com/lib/pq
```

another package testing for logs

```bash
go get github.com/stretchr/testify
```

### step 4 github actions

Refer to ci.yml to check how we setup a testing pipeline, with a postgres service, and curling the necessary golang-migration bin.

### step 5 load config using Viper

install viper

```bash
go get github.com/spf13/viper
```

create an app.env file and util/config.go to unmarshal the env variables

### step 6 create http server

use gin server
create APIs

### step 7 mock db for testing http api using stubs gomock

install gomock

```bash
go install github.com/golang/mock/mockgen@v1.6.0
```

Make sure mockgen is accessible in cli

```bash
vi ~/.zshrc
```

add the following then save and exit.

```bash
export PATH=$PATH:~/go/bin
source ~/.zshrc
```

We can mark go interfaces we'd like to mock by

```go
//go:generate mockgen -package=mockdb -destination=../mock/mock_store.go . Store
```

We run the go generate on the current project

```bash
go generate -v ./...
```

### step 8 install jwt

```bash
go get github.com/dgrijalva/jwt-go
```

install paseto

```bash
github.com/o1egl/paseto
```

## database diagram

![Alt text](schema/schema.png?raw=true "Title")
