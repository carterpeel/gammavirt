package helpers

import (
	"crypto/rand"
	"golang.org/x/sys/unix"
	"os"
)

// DirHasEnoughSpace checks to see if `dir` has enough space to add `additional` bytes to the directory
func DirHasEnoughSpace(dir string, additional int64) bool {
	var stat unix.Statfs_t
	if err := unix.Statfs(dir, &stat); err != nil {
		return false
	}
	return stat.Bavail*uint64(stat.Bsize) > uint64(additional)
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}


// NewSecureToken returns a 64 byte AES key.
func NewSecureToken() ([]byte, error) {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}