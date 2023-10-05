package utils

import (
	"fmt"
	"time"

	"github.com/NBN23dev/go-service-template/src/plugins/logger"
)

// ExecutionTime is an utility func that prints a log with the execution time of a piece of code.
func ExecutionTime(start time.Time, name string) {
	elapsed := time.Since(start)

	logger.Info(fmt.Sprintf("%s took %s", name, elapsed), nil)
}
