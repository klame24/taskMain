service-run:
	@go run cmd/main.go

up:
	docker compose up -d

down:
	docker-compose down

logs:
	docker-compose logs -f mongodb

clean:
	docker-compose down -v
	docker system prune -f

mongo-shell:
	docker exec -it taskmain-mongodb mongosh -u admin -p your_password