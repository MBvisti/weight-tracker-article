package api

type NewUserRequest struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Height int `json:"height"`
	Sex string `json:"sex"`
	ActivityLevel int `json:"activity_level"`
	Email string `json:"email"`
}

type UpdateActivityLevelRequest struct {
	Email string `json:"email"`
	ActivityLevel int `json:"activity_level"`
}