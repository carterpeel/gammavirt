package volumes

import (
	"../helpers"
	"../linker"
	"github.com/GammaByte-xyz/go-qemu"
	"github.com/google/uuid"
	libvirtxml "libvirt.org/libvirt-go-xml"
)


func NewEncryptedVolume(path string, size int64, encryptionKey []byte) (*linker.Volume, error){
	if !helpers.DirHasEnoughSpace(path, size) {
		return nil, ErrVolumeTooLarge
	}
	img, err := qemu.NewEncryptedImage(path, qemu.ImageFormatQCOW2, string(encryptionKey), uint64(size))
	if err != nil {
		return nil, err
	}
	UUID := uuid.New()
	encXml := libvirtxml.Secret{
		Ephemeral:   "no",
		Private:     "no",
		Description: UUID.String(),
		UUID:        UUID.String(),
		Usage: &libvirtxml.SecretUsage{
			Type:   "volume",
			Volume: path,
		},
	}
	EncXML, err := encXml.Marshal()
	if err != nil {
		return nil, err
	}
	vol := &linker.Volume{
		UUID:          UUID.String(),
		Path:          img.Path,
		Capacity:      size,
		EncryptionKey: encryptionKey,
		EncryptionXML: EncXML,
	}
	return vol, nil
}
