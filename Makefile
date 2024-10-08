tailwind:
	@tailwindcss -i ui/static/css/main.css -o ui/static/css/styles.css --watch

templ:
	@templ generate -watch -proxy=http://localhost:8080

air:
	@air

run:
	$(MAKE) tailwind & $(MAKE) air

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
