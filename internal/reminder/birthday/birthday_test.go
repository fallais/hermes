package birthday

import (
	"testing"
	"time"

	"hermes/internal/models"
)

func TestBirthday(t *testing.T) {
	c := models.Contact{
		Birthdate: time.Date(2010, 11, 17, 0, 0, 0, 0, time.UTC),
	}
	gb := New(true, "after", c, nil)

	if gb.GetCRONExpression() != "30 10 17 11 *" {
		t.Errorf("should be [30 10 17 11 *] but it is [%s]", gb.GetCRONExpression())
		t.Fail()
	}
}
