db_url:=postgresql://open_desk_api:postgres@localhost:5432/open_desk?sslmode=disable
init_db:
	docker run --name open_desk_db -p 5432:5432 -e POSTGRES_USER=open_desk_api -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=open_desk -d postgres

db_start:
	docker start open_desk_db

db_stop:
	docker stop open_desk_db

db_migrate_up:
	migrate -path ./db/migration -database $(db_url) up

db_migrate_down:
	migrate -path ./db/migration -database $(db_url) down