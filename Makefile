.PHONY: dev-env dev-backend dev-frontend stop clean db-cli

# Автоматично зчитуємо та експортуємо змінні з .env, якщо він існує
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Запустити інфраструктуру та чекати, поки Postgres повністю підніметься
dev-env:
	docker compose up -d
	@echo "Чекаємо на повну готовність бази даних..."
	docker compose exec -T postgres sh -c "until pg_isready -U \$${POSTGRES_USER} -d \$${POSTGRES_DB}; do sleep 0.5; done"
	@echo "PostgreSQL готовий до роботи!"

# Зупинити Postgres
stop:
	docker compose down

# Очистити дані бази даних
clean:
	docker compose down -v

# Запустити бекенд у режимі розробки (тепер команда чиста, без хардкоду зміних)
dev-backend: dev-env
	cd backend && go build -o devtrack ./cmd/app/ && ./devtrack

# Запустити фронтенд
dev-frontend:
	cd frontend && npm run dev

# Відкрити інтерактивну консоль PostgreSQL
db-cli:
	@docker compose exec -it postgres psql -U ${POSTGRES_USER} -d ${POSTGRES_DB}


# Зверніть увагу на конструкцію \$${POSTGRES_USER} у таргеті dev-env.
# Подвійний знак долара та зворотний слеш потрібні для того, щоб make не намагався підставити змінну сам,
#  а передав її безпосередньо всередину контейнера Docker.
#   А от у таргеті db-cli ми використовуємо ${POSTGRES_USER} з одним доларом, 
#   бо підставляємо її прямо з нашого зчитаного .env.
