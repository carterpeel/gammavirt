package linker

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	libvirtxml "github.com/libvirt/libvirt-go-xml"
)

func Connect(server, port, username, password string) (*Linker, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/gammavirt", username, password, server, port))
	if err != nil {
		return nil, err
	}
	var linker Linker
	linker.DB = db
	return &linker, nil
}

func (link *Linker) StoreUserData(data *UserData) error {
	_, err := link.DB.Exec("INSERT INTO users(uuid, email, password, maxRamBytes, maxVolumeBytes, maxCpuCores) VALUES (?, ?, ?, ?, ?, ?)", data.UUID, data.Email, string(data.Password), data.MaxRamBytes, data.MaxVolumeBytes, data.MaxCpuCores)
	return err
}

func (uData *UserData) CheckResources(cpuCount int64, volumeSize int64, ramSize int64) bool {
	if uData.UsedCpuCores+cpuCount > uData.MaxCpuCores || uData.UsedRamBytes+ramSize > uData.MaxRamBytes || uData.UsedVolumeBytes+volumeSize > uData.MaxVolumeBytes {
		return false
	}
	return true
}

func (dom *Domain) GenerateXML() (string, error) {
	dxml := &libvirtxml.Domain{
		Type:         "kvm",
		Name:          dom.InternalUUID,
		UUID:          dom.InternalUUID,
		Title:         dom.InternalUUID,
		MaximumMemory: &libvirtxml.DomainMaxMemory{
			Value: uint(dom.RamBytes),
			Unit:  "bytes",
		},
		Memory: &libvirtxml.DomainMemory{
			Value:    uint(dom.RamBytes),
			Unit:     "bytes",
		},
		CurrentMemory: &libvirtxml.DomainCurrentMemory{
			Value: uint(dom.RamBytes),
			Unit:  "bytes",
		},
		CPU: &libvirtxml.DomainCPU{
			Topology: &libvirtxml.DomainCPUTopology{
				Sockets: 1,
				Cores:   dom.CpuCores,
				Threads: 1,
			},
		},
		Devices: &libvirtxml.DomainDeviceList{
			Disks: []libvirtxml.DomainDisk{
				{
					Driver: &libvirtxml.DomainDiskDriver{
						Name: "qemu",
						Type: "qcow2",
						Queues: (*uint)(uint(4)),
					},
					Source: &libvirtxml.DomainDiskSource{
						File: &libvirtxml.DomainDiskSourceFile{
							File:     dom.VolPath,
						},
						Encryption: &libvirtxml.DomainDiskEncryption{
							Format: "luks",
							Secret: &libvirtxml.DomainDiskSecret{
								Type:  "passphrase",
								UUID:  dom.VolEncUUID,
							},
						},
					},
				},
			},
		},
	}
	return dxml.Marshal()
}

func (link *Linker) StoreDomainData(data *Domain) error {
	if _, err := link.DB.Exec("INSERT INTO domains(uuid, domainOwner, volPath, cpuCores, ramBytes, volumeBytes, ipAddress, hostBinding) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", data.InternalUUID, data.Owner, data.VolPath, data.CpuCores, data.RamBytes, data.VolumeBytes, data.IpAddr, data.HostBinding); err != nil {
		return err
	}
	if _, err := link.DB.Exec("UPDATE users SET usedRamBytes = usedRamBytes+?, usedCpuCores = usedCpuCores+?, usedVolumeBytes = usedVolumeBytes+? WHERE uuid = ?", data.RamBytes, data.CpuCores, data.VolumeBytes, data.Owner); err != nil {
		return err
	}
	return nil
}

func (link *Linker) StoreVolumeData(data *Volume) error {
	_, err := link.DB.Exec("INSERT INTO volumes(uuid, path, capacity, encryptionKey) VALUES (?, ?, ?, ?)", data.UUID, data.Path, data.Capacity, data.EncryptionKey)
	return err
}

func (link *Linker) GetVolumeInfoFromUUID(volumeUUID uuid.UUID) (*Volume, error) {
	vol := Volume{}
	if err := link.DB.QueryRow("SELECT * FROM volumes WHERE uuid = ?", volumeUUID.String()).Scan(&vol); err != nil {
		return nil, err
	}
	if vol.UUID == "" {
		return nil, ErrVolumeNotFound
	}
	return &vol, nil
}

func (link *Linker) GetDomainDataFromUUID(domainUUID uuid.UUID) (*Domain, error) {
	dt := Domain{}
	if err := link.DB.QueryRow("SELECT * FROM domains WHERE uuid = ?", domainUUID.String()).Scan(&dt); err != nil {
		return nil, err
	}
	if dt.InternalUUID == "" {
		return nil, ErrDomainNotFound
	}
	return &dt, nil
}

func (link *Linker) GetUserDataFromUUID(userUUID uuid.UUID) (*UserData, error) {
	var uData UserData
	if err := link.DB.QueryRow("SELECT * from users WHERE uuid = ?", userUUID.String()).Scan(&uData); err != nil {
		return nil, err
	}
	return &uData, nil
}

func (link *Linker) GetUserDataFromToken(token string) (*UserData, error) {
	var uData UserData
	if err := link.DB.QueryRow("SELECT * from users WHERE token = ?", token).Scan(&uData); err != nil {
		return nil, err
	}
	return &uData, nil
}
