package users

import (
	datetimeutils "user-profiles/datetime-utils"
)

// UserRequest represents data about a user as sent as a JSON payload in the request
type UserRequest struct {
	ID          string `json:"user_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	CreatedOn   int64  `json:"created_on"`
}

// UserResponse represents data about a user as sent back in the response
type UserResponse struct {
	ID             string `json:"user_id"`
	Name           string `json:"name"`
	DOBWeekday     string `json:"date_of_birth_weekday"`
	CreatedDateEST string `json:"created_date_est"`
}

// users slice to seed user data.
var users []UserRequest

func ConvertUserRequestsToUserResponses(userRequests []UserRequest) []UserResponse {
	var convertedUsers []UserResponse
	for _, userRequest := range userRequests {
		convertedUser := UserResponse{
			ID:             userRequest.ID,
			Name:           userRequest.Name,
			DOBWeekday:     datetimeutils.FindDayOfWeek(datetimeutils.ConvertStringToTime(userRequest.DateOfBirth)),
			CreatedDateEST: datetimeutils.ConvertUnixTimeToRFC3339(userRequest.CreatedOn),
		}
		convertedUsers = append(convertedUsers, convertedUser)
	}
	return convertedUsers
}

// GetUsers responds with the list of all users as JSON.
func GetUsers() []UserResponse {
	return ConvertUserRequestsToUserResponses(users)
}

func AddUsers(newUsers []UserRequest) []UserResponse {
	// Add the new users to the slice.
	users = append(users, newUsers...)

	// Convert the new users and return them
	return ConvertUserRequestsToUserResponses(newUsers)
}
