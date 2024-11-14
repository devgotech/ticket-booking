# createdb:
# 	docker exec -it ticket-booking-db createdb --username=postgres --owner=postgres postgres

# dropdb:
# 	docker exec -it ticket-booking-db dropdb postgres

# migratedown:
# 	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

start: 
	docker compose up --build

stop:
	docker-compose rm -v --force --stop
	docker rmi ticket-booking

PHONY: start stop