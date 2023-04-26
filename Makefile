migrate:
		migrate -path ./migrations -database 'postgres://admin:123@localhost:5432/skalm?sslmode=disable' up

dropTables:
		migrate -path ./migrations -database 'postgres://admin:123@localhost:5432/skalm?sslmode=disable' down		
