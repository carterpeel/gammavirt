package common

import (
	"../volumes"
	"github.com/google/uuid"
	"net"
	"net/mail"
	"time"
)

type Domain struct {
	InternalUUID      uuid.UUID
	RawInternalUUID   string             `json:"internal_uuid,omitempty" sql:"uuid,omitempty"`
	Owner             string             `json:"main_owner,omitempty" sql:"domainOwner,omitempty"`
	HostKernelID      string             `json:"host_kernel_id,omitempty" sql:"hostBind,omitempty"`
	NameSpace         *domainNameSpace   `json:"name_space,omitempty"`
	Users             []DomainOwner      `json:"users,omitempty"`
	HostBinding       *domainHostBinding `json:"host_binding,omitempty"`
	Network           *domainNetworkInfo `json:"network,omitempty"`
	Volume            *volumes.Volume    `json:"volume,omitempty"`
}

type domainNetworkInfo struct {
	Subnet    *net.IPNet
	IP        *net.IP
	RawIP     string     `json:"raw_ip,omitempty"`
	RawSubnet string     `json:"raw_subnet,omitempty"`
	IsWAN     bool       `json:"is_wan,omitempty"`
}

type DomainOwner struct {
	UUID            string                `json:"UUID,omitempty"`
	Auth            *auth                 `json:"auth,omitempty"`
	EmailAddress    mail.Address
	RawEmailAddress string                `json:"email_address,omitempty"`
	PermLevel       *domainOwnerPermLevel `json:"perm_level,omitempty"`
}

type auth struct {
	Token          string      `json:"token,omitempty"`
	ExpirationDate *time.Time  `json:"expiration_date,omitempty"`
}

type domainNameSpace struct {
	FullyQualifiedName string `json:"fully_qualified_name,omitempty"`
	UserChosenName     string `json:"user_chosen_name,omitempty"`
	AvailabilityZone   string `json:"availability_zone,omitempty"`
}

type domainOwnerPermLevel struct {
	Owner     bool `json:"owner,omitempty"`
	Admin     bool `json:"admin,omitempty"`
	Viewer    bool `json:"viewer,omitempty"`
}

type domainHostBinding struct {
	HostID    string  `json:"host_id,omitempty"`
	Hostname  string  `json:"hostname,omitempty"`
}

type newDomainJson struct {
	RamSize      int64         `json:"ram_size,omitempty"`
	VolumeSize   int64         `json:"volume_size,omitempty"`
	ClusterSize  int64         `json:"cluster_size,omitempty"`
	CpuCount     int64         `json:"cpu_count,omitempty"`
	AuthToken    string        `json:"auth_token,omitempty"`
	DomainName   string        `json:"domain_name,omitempty"`
	DomainIP     string        `json:"domain_ip,omitempty"`
	RandomIP     bool          `json:"random_ip,omitempty"`
	DomainOwner  DomainOwner   `json:"domain_owner,omitempty"`
}