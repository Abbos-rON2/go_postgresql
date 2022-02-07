swag-init:
	swag init -g api/main.go -o api/docs

migrate-up:
	migrate -path ./storage/migrations -database 'postgres://postgres:postgres@localhost:5432/lms?sslmode=disable' up

migrate-down:
	migrate -path ./storage/migrations -database 'postgres://postgres:postgres@localhost:5432/lms?sslmode=disable' down

migration:
	migrate create -ext sql -dir ./migrations -seq -digits 2 ${name}

sqlc:
	sqlc generate