# Application Documentation

## Requirement
1. Go version 1.20 above
2. Docker for unit test
3. Migrate, for database migration. https://github.com/golang-migrate/migrate
4. PostgresSQL Database
5. Mockery for mock in unit test, https://github.com/vektra/mockery

## Setup
1. Git clone this repository
2. Install application with command : 
```bash 
make init-app 
```
3. Copy appplication config with command : 
```bash 
cp config-example.toml config.toml
4. Edit `config.toml` file based on your own configuration.
5. Execute migration with command : 
```bash 
migrate -path migrations -database "postgresql://username:password@host:port/databasename?sslmode=disable" -verbose up
```

## Run Application
1. Run with command : 
```bash 
make run
```

