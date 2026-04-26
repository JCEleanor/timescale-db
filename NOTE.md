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

### Post-fix Typing

1. Simple Comparison
   - `int age = 25;`: Integer called age is 25
   - `var age int = 25`: Variable age, which is an int, is 25

2. Why Go does this: The "Pointer" Problem
   The creators of Go (who also helped create C) realized that C's syntax gets very confusing as code gets complex.

In C, a pointer looks like this:

```c
int *p;
```

Is the `*` attached to the int or the p? (It's attached to the p).

In Go, it is much clearer:

```go
var p *int
```

You read it left-to-right: "Variable p is a Pointer to an Int."

3. Function Signatures

This is where the syntax really shines. Look at our function again:

```go
1 func handlePostMetrics(pool *pgxpool.Pool) http.HandlerFunc
```

Reading left-to-right:

- `func`: This is a function.
- `handlePostMetrics`: Its name is this.
- `(pool *pgxpool.Pool)`: It takes a pool (pointer to a Pool).
- `http.HandlerFunc`: It returns a HandlerFunc.

In C++, a function returning a function pointer would require complex nesting of parentheses that is notoriously difficult to read. Go keeps the "What is
it?" and "What does it return?" clearly separated at the end.

4. The Short Declaration (`:=`)

Because the type comes second, Go can easily omit it when it's obvious:

```go
age := 25 // Go knows it's an int
```

### Context

TODO:

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

### JSON Tags

In Go, JSON tags are metadata attached to struct fields that control how the encoding/json package handles serialization (marshaling) and deserialization (unmarshaling). They allow you to map Go struct fields—which must be exported (start with a capital letter)—to JSON keys that follow different naming conventions

```go
type WeatherMetric struct {
	Time      time.Time `json:"time"`
	Location  string    `json:"location"`
	Tempature float64   `json:"temperature"`
	Humidity  float64   `json:"humidity"`
}

```

## Run the program

1. The "Development" Way (Run without compiling)

This compiles the code into a temporary folder and runs it immediately. It is the fastest way to test your changes.

```bash
go run main.go
# or go run .
```

2. The "Production" Way (Build an executable)

This creates a permanent binary file (an .exe on Windows or a binary on Mac/Linux) that you can run anywhere without needing the Go source code.

```bash
# 1. Compile the code
go build -o iot-server

# 2. Run the resulting file
./iot-server
```

## POSTGRES

### check container status

```bash
docker ps -a --filter "name=timescaledb"
```

### create the container

```bash
docker run -d --name timescaledb \
    -p 5432:5432 \
    -v $(pwd)/data:/pgdata \
    -e PGDATA=/pgdata \
    -e POSTGRES_PASSWORD=password \
    timescale/timescaledb-ha:pg18
```

### Enter the PostgreSQL shell & table lookup

```bash
docker exec -it timescaledb psql -U postgres

SELECT * FROM weather_metrics ORDER BY time DESC;
```

### The "Enter the Shell" Method

If you prefer to enter the container and see what's going on, you can start an interactive shell:

```bash
docker exec -it timescaledb bash
```

### create table

```bash
docker exec -i timescaledb psql -U postgres < create.sql
```

### INSERT row

```bash
curl -X POST http://localhost:8080/metrics \
-H "Content-Type: application/json" \
-d '{
    "time": "2026-04-25T14:30:00Z",
    "location": "Seattle",
    "temperature": 58.5,
    "humidity": 72.0
}'
```

### other

```
   * postgres=#: This means "I am ready for a new command."
   * postgres-#: This means "You haven't finished your last command yet; I'm waiting for more."
```

If you see the -# prompt, just type a single ; and press Enter to clear the error. Then, try again with the semicolon included:
