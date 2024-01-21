run:
		go run cmd/main.go

migrate-up:
		migrate -path ./migrations -database "postgres://postgres:test@localhost:5432/effective_mobile?sslmode=disable" up 

migrate-down:
		echo y | migrate -path ./migrations -database "postgres://postgres:test@localhost:5432/effective_mobile?sslmode=disable" down