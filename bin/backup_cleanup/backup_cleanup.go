package main

import (
	"flag"
	"fmt"
	"github.com/bborbe/backup/config"
	"github.com/bborbe/backup/dto"
	"github.com/bborbe/backup/service"
	"github.com/bborbe/backup/util"
	"github.com/bborbe/log"
	"io"
	"os"
	"sort"
)

var logger = log.DefaultLogger

func main() {
	logLevelPtr := flag.Int("loglevel", config.DEFAULT_LOG_LEVEL, "int")
	rootdirPtr := flag.String("rootdir", config.DEFAULT_ROOT_DIR, "string")
	hostPtr := flag.String("host", config.DEFAULT_HOST, "string")
	flag.Parse()
	logger.SetLevelThreshold(*logLevelPtr)
	logger.Debugf("set log level to %s", log.LogLevelToString(*logLevelPtr))

	writer := os.Stdout
	logger.Debugf("use backup dir %s", *rootdirPtr)
	backupService := service.NewBackupService(*rootdirPtr)
	err := do(writer, backupService, *hostPtr)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}
}

func do(writer io.Writer, backupService service.BackupService, hostname string) error {
	logger.Debug("start")
	var err error
	var hosts []dto.Host
	if hostname == config.DEFAULT_HOST {
		hosts, err = backupService.ListHosts()
		if err != nil {
			return err
		}
	} else {
		host, err := backupService.GetHost(hostname)
		if err != nil {
			return err
		}
		hosts = []dto.Host{host}
	}
	sort.Sort(util.HostByDate(hosts))
	for _, host := range hosts {
		err := backupService.Cleanup(host)
		if err != nil {
			return err
		}
		fmt.Fprintf(writer, "%s cleaned\n", host.GetName())
	}
	logger.Debug("done")
	return nil
}
