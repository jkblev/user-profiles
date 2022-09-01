package users

import (
	"testing"
)

// TestConvertUserRequestToUserResponse is an integration test that verifies
// that the UserRequest struct is converted/translated to a UserResponse
// with appropriate values. The underlying datetime_utils are tested further
// in user-profiles/datetime-utils_test.go
func TestConvertUserRequestsToUserResponses(t *testing.T) {
	testUsers := []UserRequest{{
		ID:          "1",
		Name:        "John Doe",
		DateOfBirth: "1983-05-12",
		CreatedOn:   0,
	}}
	received := ConvertUserRequestsToUserResponses(testUsers)
	expected := []UserResponse{{
		ID:             "1",
		Name:           "John Doe",
		DOBWeekday:     "Thursday",
		CreatedDateEST: "1969-12-31T19:00:00-05:00",
	}}

	if received[0] != expected[0] {
		t.Errorf("ConvertUserRequestToUserResponse = %#v, expected %#v", received, expected)
	}
}

// TestGetUsersEmpty verifies that we return an empty slice if no data
// has been added to the global users slice yet
func TestGetUsersEmpty(t *testing.T) {
	received := GetUsers()
	if len(received) > 0 {
		t.Errorf("GetUsers() = %#v, expected empty slice", received)
	}
}

// GetGetUsersSeeded verifies that we return the expected global
// users slice when it is seeded with preexisting data.
func TestGetUsersSeeded(t *testing.T) {
	// Seed users with an existing user to find
	existingUser := UserRequest{
		ID:          "1",
		Name:        "John Doe",
		DateOfBirth: "1983-05-12",
		CreatedOn:   0,
	}
	users = append(users, existingUser)
	received := GetUsers()
	if received[0].ID != "1" {
		t.Errorf("GetUsers() = %#v, expected %#v", received, existingUser)
	}
}

// TestAddUsers verifies that adding a new UserRequest to users (that has a
// pre-existing UserRequest already in it) will translate the new user to
// a UserResponse and only return the new user instead of all users in the
// slice.
func TestAddUsers(t *testing.T) {
	// Setup - add a test user that's already present in the users slice
	// so that we can verify the response does not include the preexisting
	// test user
	preexistingUsers := []UserRequest{{
		ID:          "1",
		Name:        "John Doe",
		DateOfBirth: "1983-05-12",
		CreatedOn:   0,
	}}
	AddUsers(preexistingUsers)

	newUsers := []UserRequest{{
		ID:          "2",
		Name:        "Jane Doe",
		DateOfBirth: "1989-04-29",
		CreatedOn:   0,
	}}

	expected := []UserResponse{{
		ID:             "2",
		Name:           "Jane Doe",
		DOBWeekday:     "Saturday",
		CreatedDateEST: "1969-12-31T19:00:00-05:00",
	}}
	received := AddUsers(newUsers)
	if received[0] != expected[0] {
		t.Errorf("ConvertUserRequestToUserResponse = %#v, expected %#v", received, expected)
	}
}
