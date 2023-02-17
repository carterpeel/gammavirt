package volumes

import "errors"

var (
	ErrVolumeTooLarge = errors.New("the requested volume is either too large to be allocated or the directory does not have enough space left")
)