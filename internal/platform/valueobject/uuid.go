package valueobject

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidUUID = errors.New("invalid UUID")

// UUID represent a unique identifier.
type UUID struct {
	value string
}

// NewUUID instantiate the VO for UUID.
func NewUUID(value string) (UUID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return UUID{}, fmt.Errorf("%w: %s", ErrInvalidUUID, value)
	}

	return UUID{
		value: v.String(),
	}, nil
}

// String type converts the UUID into string.
func (id UUID) String() string {
	return id.value
}
