package utils

import "strings"

// Constants for all Status
const (
	REVIEWING = "REVIEWING"
	ACCETPED  = "ACCEPTED"
	DENIED    = "DENIED"
)

// IsSupportedStatus returns true if the status is supported
func IsSupportedStatus(status string) bool {
	switch strings.ToUpper(status) {
	case REVIEWING, ACCETPED, DENIED:
		return true
	}
	return false
}
