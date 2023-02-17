package common

import (
	"../linker"
	"github.com/libvirt/libvirt-go"
	"github.com/sirupsen/logrus"
)

type hypervisor struct {
	Linker          *linker.Linker
	ConfigFile      *configFileData
	HypervisorInfo  *hypervisorData
	Conn            *libvirt.Connect
	Log             *logrus.Logger
}


type hypervisorData struct {
	HostID   string
	Hostname string
}

type configFileData struct {
	VolumeConfig  *volumeConfig       `yaml:"volume_config"`
	HostAlias     string              `yaml:"host_alias"`
	Listener      *hypervisorListener `yaml:"listen_config"`
	Loggers       GammaLogger         `yaml:"logger"`
}

type volumeConfig struct {
	VolumePath   string              `yaml:"volume_path"`
	VolumeFormat string              `yaml:"volume_format"`
}


type hypervisorListener struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}