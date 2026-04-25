## GO package management

- `go mod init <name>`: Initializes a new module in the current directory.
- `go get <package>@<version>` : Adds a specific dependency to your project. Use @latest for the most recent version.
- `go mod tidy`: Automatically adds missing module requirements and removes unused ones.
- `go install <package>@latest` : Downloads and installs an executable command globally.
- `go list -m all` : Lists the current module and all its dependencies.

## Go Syntax

### Short Declaration

`:=` Syntax: This is the "short variable declaration." In Go, you don't have to say `string connStr = ....` Go infers the type (it sees the quotes and knows it is a string).

- syntactic sugar `var connStr string`
- Note: You can only use `:=` inside functions.

### Context

`context.Background()`: Think of it as an "envelope" that travels with your request. It can carry deadlines or signals to stop work.

### `if err != nil`

Go does not have try/catch blocks.

### `defer`

```
pool, err := pgxpool.NewWithConfig(context.Background(), config)
// ... check error ...
defer pool.Close()
```

- defer: This tells Go: "Run this line of code at the very end of this function, right before it returns."
- Why use it?: It ensures the database connection is closed properly even if the program crashes or finishes early. It’s like setting an alarm to "clean up" later.

## Run the program

1. The "Development" Way (Run without compiling)

This compiles the code into a temporary folder and runs it immediately. It is the fastest way to test your changes.

```
go run main.go
```

2. The "Production" Way (Build an executable)

This creates a permanent binary file (an .exe on Windows or a binary on Mac/Linux) that you can run anywhere without needing the Go source code.

```
# 1. Compile the code
go build -o iot-server

# 2. Run the resulting file
./iot-server
```

## POSTGRES

### check container status

```
docker ps -a --filter "name=timescaledb"
```

### create the container

```
    docker run -d --name timescaledb \
      -p 5432:5432 \
      -v $(pwd)/data:/pgdata \
      -e PGDATA=/pgdata \
      -e POSTGRES_PASSWORD=password \
      timescale/timescaledb-ha:pg18
```

2. Execution Guide
   Once your Docker container is running, execute the following command in your terminal to initialize the schema in one go:

1 docker exec -it timescaledb psql -U postgres -c "
2 CREATE TABLE IF NOT EXISTS weather_metrics (
3 time TIMESTAMPTZ NOT NULL,
4 location TEXT NOT NULL,
5 temperature DOUBLE PRECISION NULL,
6 humidity DOUBLE PRECISION NULL
7 );
8 SELECT create_hypertable('weather_metrics', 'time', if_not_exists => TRUE);
9 "

3. Verification
   To verify that the hypertable was created successfully, you can run:

1 docker exec -it timescaledb psql -U postgres -c "\d weather_metrics"
You should see "Triggers" or "Child tables" (depending on the Postgres version) indicating that TimescaleDB is managing the table as a hypertable.

### Enter the PostgreSQL shell

```
docker exec -it timescaledb psql -U postgres
```

### The "Enter the Shell" Method

If you prefer to enter the container and see what's going on, you can start an interactive shell:

```
docker exec -it timescaledb bash
```

### create table

```
docker exec -i timescaledb psql -U postgres < create.sql
```

### other

```
   * postgres=#: This means "I am ready for a new command."
   * postgres-#: This means "You haven't finished your last command yet; I'm waiting for more."
```

If you see the -# prompt, just type a single ; and press Enter to clear the error. Then, try again with the semicolon included:
