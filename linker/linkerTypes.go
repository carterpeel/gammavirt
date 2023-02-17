package linker

import (
	"database/sql"
	"errors"
)

type Linker struct {
	DB        *sql.DB
}

type UserData struct {
	UUID            string        `sql:"uuid,omitempty"`
	Email           string        `sql:"email,omitempty"`
	Password        []byte        `sql:"password,omitempty"`
	MaxRamBytes     int64         `sql:"maxRamBytes,omitempty"`
	MaxVolumeBytes  int64         `sql:"maxVolumeBytes,omitempty"`
	MaxCpuCores     int64         `sql:"maxCpuCores,omitempty"`
	UsedRamBytes    int64         `sql:"usedRamBytes,omitempty"`
	UsedVolumeBytes int64         `sql:"usedVolumeBytes,omitempty"`
	UsedCpuCores    int64         `sql:"usedCpuCores,omitempty"`
}

type Domain struct {
	InternalUUID string `sql:"uuid,omitempty"`
	Owner        string `sql:"domainOwner,omitempty"`
	VolPath      string `sql:"volPath,omitempty"`
	VolEncUUID   string `sql:"volEncUUID,omitempty"`
	CpuCores     int    `sql:"cpuCores,omitempty"`
	RamBytes     int64  `sql:"RamBytes,omitempty"`
	VolumeBytes  int64  `sql:"VolumeBytes,omitempty"`
	IpAddr       string `sql:"ipAddress,omitempty"`
	HostBinding  string `sql:"hostBinding,omitempty"`
}

type Volume struct {
	UUID          string `sql:"uuid,omitempty"`
	Path          string `sql:"path,omitempty"`
	Capacity      int64  `sql:"capacity,omitempty"`
	EncryptionKey []byte `sql:"encryptionKey,omitempty"`
	EncryptionXML string
}

var (
	ErrVolumeNotFound = errors.New("volume does not exist")
	ErrUserNotExist   = errors.New("user does not exist")
	ErrDomainNotFound = errors.New("domain does not exist")
)