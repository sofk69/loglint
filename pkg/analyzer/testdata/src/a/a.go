package a

import (
	"fmt"
	"log/slog"
)

func test() {
	slog.Info("Starting server") // want "log message should start with lowercase"

	slog.Info("запуск сервера") // want "log message must be english" "log message contains special characters"

	slog.Info("server started!") // want "log message must be english" "log message contains special characters"

	password := "123"
	slog.Info("user password: " + password) // want "log message must be english" "log message contains special characters" "log message may contain sensitive data"

	slog.Error("Failed to connect") // want "log message should start with lowercase"

	slog.Warn("something went wrong...") // want "log message must be english" "log message contains special characters"

	slog.Info(fmt.Sprintf("Starting server")) // want "log message should start with lowercase"
}
