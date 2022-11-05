package helper

import "testing"

func TestValidateUserInput(t *testing.T) {
	result := 1 + 1

	if result != 2 {
		t.Errorf("Addition FAILED. Expected %d, got %d\n", 2, result)
	} else {
		t.Logf("Addition PASSED. Expected %d, got %d\n", 2, result)
	}

	name, mail, tickets := ValidateUserInput("abcd", "efgh", "bla@bla.com", 10, 50)

	if !name {
		t.Errorf("Name FAILED. Expected %v, got %v\n", true, name)
	} else {
		t.Logf("Name PASSED. Expected %v, got %v\n", true, name)
	}
	if !mail {
		t.Errorf("Mail FAILED. Expected %v, got %v\n", true, mail)
	} else {
		t.Logf("Mail PASSED. Expected %v, got %v\n", true, mail)
	}
	if !tickets {
		t.Errorf("Tickets FAILED. Expected %v, got %v\n", true, tickets)
	} else {
		t.Logf("Tickets PASSED. Expected %v, got %v\n", true, tickets)
	}
}
