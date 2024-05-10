run:
	go run main.go

init-app:
	go mod init user-service
	go mod tidy

docker-restart:
	docker compose down -v
	docker compose up -d

proto-gen:
	protoc proto/*/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto --experimental_allow_proto3_optional	

test-migration-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:6432/postgres?sslmode=disable" -verbose up

mock-gen:
	mockery --dir domain/user --name UserRepoInterface --filename iuser_repo.go --output domain/user/mocks --with-expecter		

test:
	go test -p 1 --v user-service/domain/user/testcase -coverprofile cover.out -coverpkg user-service/domain/user/usecase,user-service/domain/user/repo
	go tool cover -func cover.out	

test-coverage:
	go tool cover -html cover.out	