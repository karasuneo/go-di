-include .env

up:
	docker compose build && docker compose up -d

db:
	docker exec -it $(POSTGRES_HOST) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

server:
	docker exec -it di-sample-server ash
