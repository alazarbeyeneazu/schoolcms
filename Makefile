migrate-down:
	- migrate -database cockroachdb://root@localhost:26257/defaultdb?sslmode=disable -path internal/constant/query/schemas -verbose down
migrate-up:
	- migrate -database cockroachdb://root@localhost:26257/defaultdb?sslmode=disable -path internal/constant/query/schemas -verbose up
migrate-create:
	- migrate create -ext sql -dir internal/constant/query/schemas -tz "UTC" $(args)
swagger:
	- swag init -g initiator/initiator.go
tests:
	- go test ./...  -count=1
air:
	- go install github.com/cosmtrek/air@latest
sqlc:
	- sqlc generate -f ./config/sqlc.yaml
lint:
	- golangci-lint run ./...