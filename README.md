# Maps backend 🌎

Backend for my test and experiment maps project. Written in Go.


## Database and database management 💾

This project uses a postgres database using the [golang-migrate/migrate](https://github.com/golang-migrate/migrate/tree/v4.16.2/cmd/migrate) cli for the schema migrations. 

### Usage

To create a migration called `create_users_table`, run the command

```
migrate create -ext sql -dir db/migrations -seq create_users_table
```

which will setup both an `.up.sql` as well as a `.down.sql` file. To run the actual migration, write the command

```
migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

which will run all migrations. You can also run it with a numeric argument, such as 

```
migrate -database ${POSTGRESQL_URL} -path db/migrations up 1
```

to simply run 1 migration.

Currently the `POSTGRESQL_URL` is manually set in my current `.zshrc` file to the value `postgres://serveruser:secret@localhost:5432/maps?sslmode=disable` but will be replaced with a better solution.