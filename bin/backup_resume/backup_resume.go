package main

import (
	"flag"
	"io"
	"os"

	"github.com/bborbe/backup/config"
	"github.com/bborbe/backup/service"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

func main() {
	defer logger.Close()
	logLevelPtr := flag.Int("loglevel", config.DEFAULT_LOG_LEVEL, "int")
	rootdirPtr := flag.String("rootdir", config.DEFAULT_ROOT_DIR, "string")
	hostPtr := flag.String("host", config.DEFAULT_HOST, "string")
	flag.Parse()
	logger.SetLevelThreshold(*logLevelPtr)
	logger.Debugf("set log level to %s", log.LogLevelToString(*logLevelPtr))

	writer := os.Stdout
	logger.Debugf("use backup dir %s", *rootdirPtr)
	backupService := service.NewBackupService(*rootdirPtr)
	err := do(writer, backupService, *rootdirPtr, *hostPtr)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
}

func do(writer io.Writer, backupService service.BackupService, rootdirName string, hostname string) error {
	logger.Debug("start")

	logger.Debug("done")
	return nil
}