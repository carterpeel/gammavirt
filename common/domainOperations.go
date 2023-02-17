package common

import (
	"../helpers"
	"../linker"
	"../volumes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"path/filepath"
)


func (conf *hypervisor) NewDomain(r io.Reader) (*linker.Domain, error) {
	var dj *newDomainJson
	if err := json.NewDecoder(r).Decode(&dj); err != nil {
		return nil, err
	}
	switch {
	case dj.AuthToken == "":
		return nil, ErrTokenMissing
	case len(dj.AuthToken) < 24:
		return nil, ErrTokenMalformed
	case dj.RamSize == 0:
		return nil, ErrMissingRamSize
	case dj.VolumeSize == 0:
		return nil, ErrMissingDiskSize
	case dj.CpuCount == 0:
		return nil, ErrMissingCpuCount
	case dj.RandomIP == true && dj.DomainIP != "":
		return nil, ErrRandomAddressExclusive
	}

	userData, err := conf.Linker.GetUserDataFromToken(dj.AuthToken)
	if err != nil {
		return nil, err
	}

	if !userData.CheckResources(dj.CpuCount, dj.VolumeSize, dj.RamSize) {
		return nil, ErrNotEnoughUserResources
	}

	DomUUID := uuid.New()
	volUUID := uuid.New()
	token, err := helpers.NewSecureToken()
	if err != nil {
		return nil, err
	}
	vol, err := volumes.NewEncryptedVolume(filepath.Join(conf.ConfigFile.VolumeConfig.VolumePath, fmt.Sprintf("%s-%s.%s", volUUID.String(), dj.DomainName, conf.ConfigFile.VolumeConfig.VolumeFormat)), dj.VolumeSize, token)
	if err != nil {
		return nil, err
	}
	if err := conf.Linker.StoreVolumeData(vol); err != nil {
		return nil, err
	}
	encConf, err := conf.Conn.SecretDefineXML(vol.EncryptionXML, 0)
	if err != nil {
		return nil, err
	}
	if err := encConf.SetValue(token, 0); err != nil {
		return nil, err
	}
	nj := &linker.Domain{
		InternalUUID: DomUUID.String(),
		Owner:        userData.UUID,
		VolPath:      vol.Path,
		VolEncUUID:   vol.UUID,
		CpuCores:     int(dj.CpuCount),
		RamBytes:     dj.RamSize,
		VolumeBytes:  dj.VolumeSize,
		IpAddr:       "0.0.0.0",
		HostBinding:  conf.HypervisorInfo.HostID,
	}
	if err := conf.Linker.StoreDomainData(nj); err != nil {
		return nil, err
	}
	domainXML, err := nj.GenerateXML()
	if err != nil {
		return nil, err
	}
	dom, err := conf.Conn.DomainDefineXML(domainXML)
	if err != nil {
		return nil, err
	}
	if err := dom.SetAutostart(true); err != nil {
		return nil, err
	}
	if err := dom.Create(); err != nil {
		return nil, err
	}
	return nj, nil
}
