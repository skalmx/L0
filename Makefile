.PHONY:
.DEFAULT_GOAL := run
build:
		go mod download
		go build -o  ./.bin/app cmd/main.go

run:	build compose migrate
		./.bin/app

migrate:
		migrate -path ./migrations -database 'postgres://admin:123@localhost:5432/skalm?sslmode=disable' up

dropTables:
		migrate -path ./migrations -database 'postgres://admin:123@localhost:5432/skalm?sslmode=disable' down		

compose:
		docker-compose up -d --remove-orphans

publish:
		go build -o ./.bin/publisher cmd/publisher/pub.go
		./.bin/publisher