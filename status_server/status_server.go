package status_server

import (
	"github.com/bborbe/backup/service"
	"github.com/bborbe/backup/status_checker"
	"github.com/bborbe/backup/status_handler"
	"github.com/bborbe/server"
)

func NewServer(port int, rootdir string) server.Server {
	backupService := service.NewBackupService(rootdir)
	statusChecker := status_checker.NewStatusChecker(backupService)
	handler := status_handler.NewStatusHandler(statusChecker)
	return server.NewServerPort(port, handler)
}
