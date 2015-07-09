# Compadre
Compadre is a lightweight Go package for application configuration using
[JSON](http://json.org/). It provides a thin layer over
[encoding/json](http://golang.org/pkg/encoding/json/) allowing you to
interpolate environment variables with default values.

## Installation

1. Get it

`go get github.com/appriser/compadre`

2. Include it

`import "github.com/appriser/compadre"`

## Usage

Define your application settings:

```json
{
  "app": {
    "port": 3000
  },
  "database": {
    "timeout": "%d{DB_TIMEOUT:1000}"
    "connectionString": "dbname=testdb sslmode=#s{DB_USE_SSL:true}"
  }
}
```
Describe your configuration object using structs:

```go
type Config struct {
	App      AppConfig
	Database DatabaseConfig
}

type AppConfig struct {
	Port int
}

type DatabaseConfig struct {
	Timeout          int
	ConnectionString string
}
```

Sit back and let Compadre handle the rest for you:

```go
func main() {
	var config Config
	if err := compadre.Read("config.json", &config); err != nil {
		log.Fatalf("Failed to read configuration file: %s\n", err)
	}

	fmt.Println(config.App.Port)                  // => 3000
	fmt.Println(config.Database.Timeout)          // => 1000
	fmt.Println(config.Database.ConnectionString) // => "dbname=testdb sslmode=true"
}
```

## Details

Compadre empowers your standard JSON file by allowing you to define expressions
that expand during runtime.

Expressions can be of type: int (`#d`), bool (`#t`) and string (`#s`) and have
a default value, but it's not required. If a given environment variable is
empty and no default value is set, it will use the type default which is `0`
for int, `false` for bool and `""` for string.
