FROM golang:1.22-alpine
WORKDIR /app

# Устанавливаем необходимые пакеты
RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add make

# Копируем и устанавливаем зависимости go проекта
COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

# Для просмотра изменений в go файлах установим демон Compile Daemon
# RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon

# Копируем все файлы и запускаем сборку
COPY . .

COPY ./entrypoint.sh /entrypoint.sh

# для работы с wait-for требуется bash, который по умолчанию не поставляется с alpine. Вместо этого используйте wait-for
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh", "/entrypoint.sh" ]
