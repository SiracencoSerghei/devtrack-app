.PHONY: dev-env dev-backend dev-frontend stop clean

# Запустити інфраструктуру та чекати, поки Postgres повністю підніметься
dev-env:
	docker compose up -d
	@echo "Чекаємо на повну готовність бази даних..."
	docker compose exec -T postgres sh -c "until pg_isready -U postgres -d devtrack; do sleep 0.5; done"
	@echo "PostgreSQL готовий до роботи!"

# Зупинити Postgres
stop:
	docker compose down

# Очистити дані бази даних
clean:
	docker compose down -v

# Запустити бекенд у режимі розробки
dev-backend: dev-env
	cd backend && go build -o devtrack ./cmd/app/ && ./devtrack

# Запустити фронтенд
dev-frontend:
	cd frontend && npm run dev

# Відкрити інтерактивну консоль PostgreSQL всередині контейнера
db-cli:
	@docker compose exec -it postgres psql -U postgres -d devtrack