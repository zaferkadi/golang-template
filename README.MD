### Step 1

Install golang-migrate to start creating migration files

```bash
brew install golang-migrate
```

Create the first sql migration file

```bash
migrate create -ext sql -dir db/migration -seq init_schema
```

Add your sql schemas into up.sql and down.sql files inside db/migration folder

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
