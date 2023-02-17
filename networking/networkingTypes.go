package networking

import (
	"../common"
	"github.com/google/uuid"
	"net"
)

type UserNetwork struct {
	User      *common.DomainOwner
	UUID      *uuid.UUID
	RawSubnet string `json:"subnet"`
	Subnet    *net.IPNet
	Owners    *common.DomainOwner
}

