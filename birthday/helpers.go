package birthday

import (
	"time"

	"gobirthday/models"
)

// calculateAge returns the age based on the birthdate. 0 if there is no year.
func calculateAge(birthdate models.Birthdate) int {
	if birthdate.Year == 0 {
		return 0
	}

	return time.Now().Year()-birthdate.Year
}