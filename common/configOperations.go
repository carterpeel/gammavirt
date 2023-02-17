package common

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func ReadConfig() (*configFileData, error) {
	var cf configFileData
	confBytes, err := ioutil.ReadFile("/etc/gammavirt/config.yml")
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(confBytes, &cf); err != nil {
		return nil, err
	}
	return &cf, nil
}

func (conf *hypervisor) ReloadConfig() error {
	yamlBytes, err := ioutil.ReadFile("/etc/gammavirt/config.yml")
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(yamlBytes, &conf.ConfigFile); err != nil {
		return err
	}
	return nil
}

func GenerateConfigTemplate() error {
	config := configFileData{
		VolumeConfig: &volumeConfig{
			VolumePath:   "/var/lib/libvirt/images",
			VolumeFormat: "qcow2",
		},
		HostAlias: "CHANGE-ME",
		Listener: &hypervisorListener{
			Address: "0.0.0.0",
			Port:    "8080",
		},
		Loggers: GammaLogger{
			RemoteSyslogs: []syslogEndpoint{
				{
					Address:  "127.0.0.1:514",
					Protocol: "UDP",
					Tag:      "GammaVirt",
				},
			},
			LogFiles:      []string{
				"/var/log/gammavirt.log",
			},
		},
	}
	yamlBytes, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile("/tmp/config-template.yml", yamlBytes, 0644)
}