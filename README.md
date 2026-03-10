loglint

Линтер для Go, который проверяет лог-сообщения на соответствие установленным правилам. Совместим с `golangci-lint`.

Правила:
1) lowercase — сообщение должно начинаться со строчной буквы.
2) english — сообщение должно быть на английском языке.
3) special-chars — сообщение не должно содержать спецсимволы и эмодзи.
4) sensitive — сообщение не должно содержать чувствительные данные: `password`, `token`, `api_key` и другие.

Структура проекта:
loglint/
├── cmd/loglint/main.go
├── pkg/analyzer/
│   ├── analyzer.go
│   ├── analyzer_test.go
│   ├── log_call.go
│   ├── plugin.go
│   ├── config/config.go
│   └── rules/
│       ├── lowercase.go
│       ├── english.go
│       ├── special_chars.go
│       └── sensitive.go
├── testdata/src/
│   ├── a/a.go
│   └── valid/valid.go
├── .github/workflows/ci.yml
├── .golangci.yml
└── go.mod


Сборка:
git clone https://github.com/sofk69/loglint
cd loglint
go mod download
go build ./...


Запуск тестов:
go test ./...

Использование:
Standalone
go run ./cmd/loglint ./...

Через golangci-lint:
Соберите плагин как `.so` файл:
go build -buildmode=plugin -tags=golangci -o bin/loglint.so .
Добавьте в `.golangci.yml`:

linters:
  enable:
    - loglint

linters-settings:
  custom:
    loglint:
      path: ./bin/loglint.so
      description: "Log message linter"
      original-url: github.com/sofk69/loglint

Запустите:
golangci-lint run ./...

Конфигурация:
Можно отключить отдельные правила или добавить свои ключевые слова для проверки чувствительных данных:
linters-settings:
  custom:
    loglint:
      path: ./bin/loglint.so
      settings:
        disable-lowercase: false
        disable-english: false
        disable-special-chars: false
        disable-sensitive: false
        extra-sensitive-keywords:
          - "jwt"
          - "bearer"
          - "access_token"

Примеры:
// Ошибки
slog.Info("Starting server")          // должно начинаться со строчной буквы
slog.Info("запуск сервера")           // только английский язык
slog.Info("server started!")          // спецсимволы запрещены
slog.Info("user password: " + pass)   // чувствительные данные

// Корректно
slog.Info("server started")
slog.Info("connected to database")
slog.Error("failed to connect")
slog.Info("user authenticated successfully")


Поддерживаемые логгеры
- `log/slog`
- `go.uber.org/zap`
