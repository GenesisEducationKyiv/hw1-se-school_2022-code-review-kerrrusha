package logger

import (
	"time"
)

type LogBuilder struct{}

func (lb *LogBuilder) buildTime(logMsg string) string {
	return time.Now().Format("2006.01.02 15:04:05") + " - " + logMsg
}

func (lb *LogBuilder) BuildDebug(msg string) string {
	return lb.buildTime("[DEBUG] " + msg)
}

func (lb *LogBuilder) BuildInfo(msg string) string {
	return lb.buildTime("[INFO] " + msg)
}

func (lb *LogBuilder) BuildError(msg string) string {
	return lb.buildTime("[ERROR] " + msg)
}
