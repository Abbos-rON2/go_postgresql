swag-init:
	swag init -g api/main.go -o api/docs

migrate-up:
	migrate -path ./migrations -database 'postgres://abbos:root@localhost:5432/go_postgresql?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://abbos:root@localhost:5432/go_postgresql?sslmode=disable' down

migration:
	migrate create -ext sql -dir ./migrations -seq -digits 2 ${name}