// +build prod

package loggermdl

// LogDebug logs a message at level Debug on the standard logger.
func LogDebug(args ...interface{}) {
	sugar.Debug(args)
}

// LogInfo logs a message at level Info on the standard logger.
func LogInfo(args ...interface{}) {
	sugar.Info(args)
}

// LogWarn logs a message at level Warn on the standard logger.
func LogWarn(args ...interface{}) {
	sugar.Warn(args)
}

// LogError logs a message at level Error on the standard logger.
func LogError(args ...interface{}) {
	sugar.Error(args)
}

// LogFatal logs a message at level Fatal on the standard logger.
func LogFatal(args ...interface{}) {
	sugar.Fatal(args)
}

// Log as an Info but highlights it.
func LogSpot(args ...interface{}) {
	sugar.Info(args)
}

// Panic logs a message at level Panic on the standard logger.
func LogPanic(args ...interface{}) {
	sugar.Panic(args)
}
