start:
	docker compose up --build -d

stop:
	docker compose down

logs:
	docker compose logs -f --tail 100 server
