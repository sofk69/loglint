package a

import (
	"fmt"
	"log/slog"
)

func test(password string, name string) {

	slog.Info("Starting server") // want "log message should start with lowercase"

	slog.Info("запуск сервера") // want "log message must be english"

	slog.Info("server started!") // want "log message contains special characters"

	slog.Info("password: " + password) // want "log message may contain sensitive data"

	slog.Info(fmt.Sprintf("User %s logged", name)) // want "log message should start with lowercase"

}
