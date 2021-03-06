package host

import (
	"errors"
	"fmt"
	"os"

	"github.com/bborbe/backup/fileutil"
	"github.com/bborbe/backup/rootdir"
	"github.com/bborbe/log"
)

type host struct {
	name    string
	rootdir rootdir.Rootdir
}

type Host interface {
	Path() string
	Name() string
}

var logger = log.DefaultLogger

func ByName(rootdir rootdir.Rootdir, name string) Host {
	h := new(host)
	h.rootdir = rootdir
	h.name = name
	return h
}

func All(root rootdir.Rootdir) ([]Host, error) {
	file, err := os.Open(root.Path())
	if err != nil {
		logger.Debugf("open rootdir %s failed: %v", root.Path(), err)
		return nil, err
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		logger.Debugf("file stat failed: %v", err)
		return nil, err
	}
	if !fileinfo.IsDir() {
		msg := fmt.Sprintf("rootdir %s is not a directory", root.Path())
		logger.Debug(msg)
		return nil, errors.New(msg)
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		logger.Debugf("read dir names failed: %v", err)
		return nil, err
	}
	hosts := make([]Host, 0)
	for _, name := range names {
		host := ByName(root, name)
		isDir, err := fileutil.IsDir(host.Path())
		if err != nil {
			return nil, err
		}
		if isDir {
			hosts = append(hosts, host)
		}
	}
	return hosts, nil
}

func (h *host) Path() string {
	return fmt.Sprintf("%s%c%s", h.rootdir.Path(), os.PathSeparator, h.name)
}

func (h *host) Name() string {
	return h.name
}
