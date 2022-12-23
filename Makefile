pi-dev:
	docker compose down && docker compose -f docker-compose.yml -f docker-compose.pi-dev.yml up -d --remove-orphans
win-dev:
	docker compose down && docker compose -f docker-compose.yml -f docker-compose.win-dev.yml up -d --remove-orphans
down:
	docker compose -f docker-compose.yml down
mongo-repair:
	docker compose -f docker-compose.pi-dev-repair.yml up
