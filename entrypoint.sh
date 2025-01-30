#!/bin/sh

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="host=postgres port=5432 user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_NAME}"

# Виконуємо міграції
echo "Running migrations..."
goose -dir /app/migrations up

# Перевірка статусу міграцій
if [ $? -ne 0 ]; then
  echo "Migrations failed!"
  exit 1
fi

# Тепер запускаємо сервер
echo "Starting application..."
./main