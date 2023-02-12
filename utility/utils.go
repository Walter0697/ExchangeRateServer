package utility

import (
	"errors"

	"gorm.io/gorm"
)

func RecordNotFound(err error) bool {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	}
	return false
}
