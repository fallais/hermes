package birthday

import (
	"testing"
	"time"

	"hermes/internal/models"
)

func TestBirthday(t *testing.T) {
	gb := New(true, "after", nil, nil, nil)

	if gb.NbContacts() != 0 {
		t.Errorf("should be 0 but it is %d", gb.NbContacts())
		t.Fail()
	}
}

func TestPrepareMessage(t *testing.T) {
	mt := map[string]string{
		"header": "hi !",
		"base":   "this is the day of {{contact}} !",
		"age":    "{{age}} years old !",
		"footer": "bye !",
	}

	c := &models.Contact{
		Birthdate: time.Now().AddDate(-1, 0, 0),
	}

	gb := New(true, "after", mt, nil, nil)

	if gb.prepareMessage(c) != "hi ! this is the day of {{contact}} ! {{age}} years old ! bye !" {
		t.Errorf(gb.prepareMessage(c), c)
		t.Fail()
	}
}
