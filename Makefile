migrate-up:
	migrate -database "postgresql://postgres:As123456@@localhost:5432/testapp?sslmode=disable" -path migrations up


migrate-down:
	migrate -database "postgresql://postgres:As123456@@localhost:5432/testapp?sslmode=disable" -path migrations down


start:
	go run .\cmd\main.go