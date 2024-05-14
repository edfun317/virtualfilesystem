package event

import (
	"iscoollab/filesystem/domain"
	"testing"
)

func TestRegister(t *testing.T) {

	// Assuming Users is initialized properly and AddUser method behavior is as expected.
	users := domain.NewUsers()
	// Prepare and inject dependencies if possible (not shown here as depends on actual implementation)

	// Success case: Adding a new user
	err := Register(users, "testuser")
	if err != nil {
		t.Errorf("Register() failed, expected success, got error: %s", err)
	}

	// Failure case: Adding a duplicate user
	// This assumes that AddUser handles and returns an error on duplicate user addition.
	// Add the user first to simulate the duplicate scenario.
	_ = users.AddUser("newuser")
	err = Register(users, "newuser")
	if err == nil {
		t.Errorf("Register() should fail on duplicate user, but it did not")
	}
}
