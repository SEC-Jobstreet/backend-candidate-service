package aggregate

import "github.com/pkg/errors"

var (
	ErrProfileAlreadyCompleted = errors.New("Profile already completed")
	ErrAlreadySubmitted        = errors.New("already submitted")
	ErrProfileNotFound         = errors.New("profile not found")
	ErrAlreadyCreated          = errors.New("profile with given id already created")
)
