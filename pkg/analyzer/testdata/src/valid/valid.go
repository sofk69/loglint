package valid

import "log/slog"

func test() {
	slog.Info("server started")
	slog.Info("connected to database")
	slog.Error("failed to connect")
	slog.Warn("retry attempt 3")
	slog.Debug("request completed")
	slog.Info("user authenticated successfully")
}
