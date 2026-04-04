FROM migrate/migrate:v4.18.3

# Устанавливаем дополнительные утилиты для работы с пользователями
RUN apk add --no-cache shadow

ARG UID=1000
ARG GID=1000
ENV UID=${UID}
ENV GID=${GID}

USER root

RUN addgroup -g ${GID} appgroup && \
    adduser -D -u ${UID} -G appgroup appuser

# Устанавливаем рабочую директорию для миграций
WORKDIR /migrations

# Меняем пользователя
USER appuser
