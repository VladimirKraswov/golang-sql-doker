wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

# Просматриваем все .go фалы и вызываем go build если в файлах будут изменения.
CompileDaemon --build="go build -o main ./cmd/app/main.go"  --command=./main