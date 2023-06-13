migration_up: migrate -path db/migration/ -database "postgresql\://postgres\:postgres@localhost\:5432/go-boiler-plate?sslmode=disable" -verbose up

migration_down: migrate -path db/migration/ -database "postgresql\://postgres\:postgres@localhost\:5432/go-boiler-plate?sslmode=disable" -verbose down

migration_fix: migrate -path db/migration/ -database "postgresql\://postgres\:postgres@localhost\:5432/go-boiler-plate?sslmode=disable" force VERSION
