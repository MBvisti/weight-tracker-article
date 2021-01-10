package api

import "time"

type User struct {
	ID            int       `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	Age           int       `json:"age"`
	Height        int       `json:"height"`
	Sex           string    `json:"sex"`
	ActivityLevel int       `json:"activity_level"`
	WeightGoal    string    `json:"weight_goal"`
	Email         string    `json:"email"`
}

type NewUserRequest struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	Height        int    `json:"height"`
	Sex           string `json:"sex"`
	ActivityLevel int    `json:"activity_level"`
	WeightGoal    string `json:"weight_goal"`
	Email         string `json:"email"`
}

type UpdateActivityLevelRequest struct {
	Email         string `json:"email"`
	ActivityLevel int    `json:"activity_level"`
}

type Weight struct {
	Weight             int `json:"weight"`
	UserID             int `json:"user_id"`
	BMR                int `json:"bmr"`
	DailyCaloricIntake int `json:"daily_caloric_intake"`
}

type NewWeightRequest struct {
	Weight int `json:"weight"`
	UserID int `json:"user_id"`
}
