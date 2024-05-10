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
```
4. Edit `config.toml` file based on your own configuration.
5. Execute migration with command : 
```bash 
migrate -path migrations -database "postgresql://username:password@host:port/databasename?sslmode=disable" -verbose up
```
6. Create protobuf folder with command : 
```bash 
mkdir pb
```
7. Generate proto with command : 
```bash 
make proto-gen
```

## Run Application
1. Run with command : 
```bash 
make run
```

## Run Unit Test
1. Run docker with command : 
```bash 
make docker-restart
```
2. Data migration for testing with command :
```bash 
make test-migration-up
```
3. Run mock with command : 
```bash 
make mock-gen
```
4. Run unit test with command : 
```bash 
make test
```
5. To view coverage report in browser : 
```bash 
make test-coverage
```