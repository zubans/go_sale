.PHONY: build run start stop restart init-db compile clean logs

# Путь к бинарному файлу
BINARY=main

# Шаги сборки Docker-контейнеров
build:
	docker-compose build

# Запуск всех контейнеров в фоновом режиме
run:
	docker-compose up -d

# Запуск только контейнера приложения
start:
	docker-compose up -d app

# Остановка и удаление всех контейнеров
stop:
	docker-compose down

# Перезапуск контейнеров
restart: stop run

# Инициализация базы данных (с удалением существующих данных) и запуск приложения
init-db:
	docker-compose down -v
	docker-compose up -d postgres_db
	sleep 2 # Ждем, пока БД полностью инициализируется
	docker-compose up -d app

# Компиляция Go-проекта
compile:
	go mod tidy
	go build -o $(BINARY)

# Очистка скомпилированных файлов и Docker контейнеров
clean:
	docker-compose down -v
	rm -f $(BINARY)

# Показ логов всех контейнеров
logs:
	docker-compose logs -f