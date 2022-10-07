package logger

import (
	"fmt"
)

type ConsoleLogger struct {
	*LogBuilder
}

func (c *ConsoleLogger) log(logMsg string) {
	fmt.Println(logMsg)
}

func (c *ConsoleLogger) Info(msg string) {
	c.log(c.BuildInfo(msg))
}

func (c *ConsoleLogger) Debug(msg string) {
	c.log(c.BuildDebug(msg))
}

func (c *ConsoleLogger) Error(msg string) {
	c.log(c.BuildError(msg))
}
