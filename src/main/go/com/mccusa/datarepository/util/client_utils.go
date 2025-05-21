package util

import "strings"

// ClientUtils holds utility methods for client operations.
type ClientUtils struct{}

// NewClientUtils constructs a new ClientUtils.
func NewClientUtils() ClientUtils {
	return ClientUtils{}
}

// ConcatenateNameParts joins non-empty name parts with spaces.
func (u ClientUtils) ConcatenateNameParts(firstName string, secondName string, lastName string, secondLastName string) string {
	parts := []string{}
	if firstName != "" {
		parts = append(parts, firstName)
	}
	if secondName != "" {
		parts = append(parts, secondName)
	}
	if lastName != "" {
		parts = append(parts, lastName)
	}
	if secondLastName != "" {
		parts = append(parts, secondLastName)
	}
	return strings.Join(parts, " ")
}
