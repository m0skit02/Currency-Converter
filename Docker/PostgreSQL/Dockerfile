# Используем официальный образ PostgreSQL
FROM postgres:latest

# Определяем переменные окружения
ENV POSTGRES_DB=converter
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=1973

# Проброс порта
EXPOSE 5432

# Определяем точку монтирования для тома
VOLUME /var/lib/postgresql/data
