package volumes

import (
	"github.com/GammaByte-xyz/go-qemu"
	"github.com/google/uuid"
)

type Volume struct {
	Size          int64
	UUID          uuid.UUID
	RawUUID       string
	Path          string
	EncryptionKey []byte
	QemuData      *qemu.Image
}
