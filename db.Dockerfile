# Возмем базовый образ
FROM mysql:8.0.23

# Импорт данных в контейнер
# Все скрипты в docker-entrypoint-initdb.d/ автоматически выполняются при запуске контейнера
COPY ./database/*.sql /docker-entrypoint-initdb.d/