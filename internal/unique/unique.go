// package unique defines generation of id's for data
package unique

import (
	"github.com/google/uuid"
)

// GenerateID provides generating of ids on UUID
func GenerateID() string {
	return guuid.New().String()
}

