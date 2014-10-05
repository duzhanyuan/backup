package rootdir

import (
	"os"
	"fmt"
	"errors"
	"github.com/bborbe/log"
)

type Rootdir string

var logger = log.DefaultLogger

func New(dir string) Rootdir {
	return Rootdir(dir)
}

func (r *Rootdir) Names() ([]string, error ) {
	file, err := os.Open(string(*r))
	if err != nil {
		logger.Debugf("open rootdir %s failed: %v", r, err)
		return nil, err
	}
	defer file.Close()
	fileinfo, err := file.Stat()
	if err != nil {
		logger.Debugf("file stat failed: %v", err)
		return nil, err
	}
	if !fileinfo.IsDir() {
		msg := fmt.Sprintf("rootdir %s is not a directory", r)
		logger.Debug(msg)
		return nil, errors.New(msg)
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		logger.Debugf("read dir names failed: %v", err)
		return nil, err
	}
	return names, nil
}


