package a

import "log/slog"

func test() {
	slog.Info("Starting server") // want "log message must start with lowercase"
	slog.Info("запуск сервера")  // want "log message must be in english"
	slog.Info("server started!") // want "log message must not contain special characters"
	password := "123"
	slog.Info("user password: " + password) // want "log message contains sensitive data"
}
