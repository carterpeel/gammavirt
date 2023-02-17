package common

import (
	"github.com/sirupsen/logrus"
	rSyslog "github.com/sirupsen/logrus/hooks/syslog"
	"io"
	"log/syslog"
	"os"
	"strings"
)

func newLogger(gl *GammaLogger) (*logrus.Logger, error) {
	Log := logrus.New()
	writers := make([]io.Writer, len(gl.LogFiles))
	for _, r := range gl.RemoteSyslogs {
		hook, err := rSyslog.NewSyslogHook(strings.ToLower(r.Protocol), r.Address, syslog.LOG_DEBUG, r.Tag)
		if err != nil {
			continue
		}
		Log.Hooks.Add(hook)
	}
	if  len(gl.LogFiles) == 0 {
		return Log, nil
	}
	for i, filepath := range gl.LogFiles {
		fi, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644);
		if err != nil {
			return Log, err
		}
		writers[i] = fi
	}
	writers = append(writers, os.Stderr)
	Log.SetNoLock()
	Log.Out = io.MultiWriter(writers...)
	return Log, nil
}