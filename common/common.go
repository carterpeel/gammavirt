package common

import (
	"../helpers"
	"../linker"
	"fmt"
	"github.com/libvirt/libvirt-go"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)


func init() {
	if !helpers.Exists("/etc/gammavirt/config.yml") {
		if err := GenerateConfigTemplate(); err != nil {
			panic(err)
		}
		fmt.Println("A template config file has been generated: '/tmp/config-template.yml'")
		fmt.Println("Please update it accordingly and move it to '/etc/gammavirt/config.yml'.")
		os.Exit(1)
	}
}

func ConnectHost() (*hypervisor, error) {
	var hv hypervisor
	var err error
	hv.ConfigFile, err = ReadConfig()
	if err != nil {
		return nil, err
	}
	if len(hv.ConfigFile.Loggers.LogFiles) == 0 && len(hv.ConfigFile.Loggers.RemoteSyslogs) == 0 {
		hv.Log = logrus.New()
		logFile, err := os.OpenFile("/var/log/gammavirt.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			return nil, err
		}
		hv.Log.Out = io.MultiWriter(os.Stderr, logFile)
		hv.Log.SetNoLock()
	} else {
		hv.Log, err = newLogger(&GammaLogger{
			RemoteSyslogs: hv.ConfigFile.Loggers.RemoteSyslogs,
			LogFiles:      hv.ConfigFile.Loggers.LogFiles,
		})
		if err != nil {
			return nil, err
		}
	}
	hv.Conn, err = libvirt.NewConnect("qemu:///system?socket=/var/run/libvirt/libvirt-sock")
	if err != nil {
		return nil, err
	}
	hv.Linker, err = linker.Connect("localhost", "3306", "root", "password")
	if err != nil {
		return nil, err
	}
	return &hv, nil
}
