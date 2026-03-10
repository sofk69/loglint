loglint

Линтер для Go, который проверяет лог-сообщения на соответствие установленным правилам. Совместим с `golangci-lint`.

Правила:
1) lowercase — сообщение должно начинаться со строчной буквы.
2) english — сообщение должно быть на английском языке.
3) special-chars — сообщение не должно содержать спецсимволы и эмодзи.
4) sensitive — сообщение не должно содержать чувствительные данные: `password`, `token`, `api_key` и другие.

Структура проекта:
cmd/loglint/main.go — точка входа для standalone запуска

pkg/analyzer/analyzer.go — основной анализатор

pkg/analyzer/analyzer_test.go — интеграционный тест

pkg/analyzer/log_call.go — поиск лог-вызовов в AST

pkg/analyzer/plugin.go — плагин для golangci-lint

pkg/analyzer/config/config.go — структура конфигурации

pkg/analyzer/rules/lowercase.go — правило: строчная буква

pkg/analyzer/rules/english.go — правило: английский язык

pkg/analyzer/rules/special_chars.go — правило: спецсимволы и эмодзи

pkg/analyzer/rules/sensitive.go — правило: чувствительные данные

testdata/src/a/a.go — тестовые примеры с ошибками

testdata/src/valid/valid.go — тестовые примеры без ошибок

.github/workflows/ci.yml — GitHub Actions CI

.golangci.yml — конфигурация golangci-lint

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
