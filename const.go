package swaggen

import "errors"

var (
	ErrInvalidVersion = errors.New("an appropriate spec could not be found")
	ErrInvalidSource  = errors.New("an appropriate source processor could not be found")
)
