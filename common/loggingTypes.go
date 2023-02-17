package common

import (
	"github.com/sirupsen/logrus"
	"io"
)

type syslogEndpoint struct {
	Address  string `yaml:"address"`
	Protocol string `yaml:"protocol"`
	Tag      string `yaml:"tag"`
}

type GammaLogger struct {
	RemoteSyslogs []syslogEndpoint `yaml:"syslog_endpoints"`
	LogFiles      []string         `yaml:"syslog_files"`
}

type VSlurp struct {
	Log           *logrus.Logger
	gammaConf     GammaLogger
	outWriter     io.Writer
}
