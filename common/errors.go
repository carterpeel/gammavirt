package common

import "errors"

var (
	// User-invoked errors
	ErrNotEnoughUserResources   = errors.New("user lacks the resources required to provision a new domain")
	ErrNotEnoughSystemResources = errors.New("system chosen lacks the resources required to provision a new domain")
	ErrMissingCpuCount          = errors.New("missing vCPU count field")
	ErrMissingRamSize           = errors.New("missing RAM size field")
	ErrMissingDiskSize          = errors.New("missing disk size field")
	ErrRandomAddressExclusive   = errors.New("the random_ip field is exclusive")

	// Authorization errors
	ErrTokenMissing             = errors.New("missing authentication token")
	ErrTokenTooShort            = errors.New("token must be 24 characters or more")
	ErrTokenMalformed           = errors.New("token is malformed (contact an administrator)")
	ErrTokenInvalid             = errors.New("token did not match any known authentication keys")

	ErrBadJsonInput             = errors.New("JSON input was either malformed or nonexistent")
)
