run:
	go run main.go

init-app:
	go mod init user-service
	go mod tidy

docker-restart:
	docker compose down -v
	docker compose up -d

proto-gen:
	protoc --proto_path=proto --go_out=paths=source_relative,plugins=grpc:./pb proto/*/*.proto
	# for support libprotoc > 3.17.*
	# protoc -I=proto/location --go_out=plugins=grpc:. proto/*/*.proto

proto-gen2:
	protoc proto/*/*.proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:pb --go-grpc_opt=paths=source_relative -I=proto --experimental_allow_proto3_optional	

test-migration-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:6432/postgres?sslmode=disable" -verbose up