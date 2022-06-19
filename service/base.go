package service

import (
	"os"
	"sync"
	"time"
)

const defaultDatabaseTimeout = "5s"

var dbDuration *time.Duration
var one sync.Once

func InitDBTimeOut() {
	one.Do(func() {
		timeoutStr := os.Getenv("DATABASE_TIMEOUT")
		if timeoutStr == "" {
			dur, _ := time.ParseDuration(defaultDatabaseTimeout)
			dbDuration = &dur
		}

		dur, _ := time.ParseDuration(timeoutStr)
		dbDuration = &dur
	})
}
