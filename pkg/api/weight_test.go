package api_test

import (
	"errors"
	"reflect"
	"testing"
	"weight-tracker/pkg/api"
)

type mockWeightRepo struct{}

func (m mockWeightRepo) CreateWeightEntry(w api.Weight) error {
	return nil
}

func (m mockWeightRepo) GetUser(userID int) (api.User, error) {
	if userID != 1 {
		return api.User{}, errors.New("storage - user doesn't exists")
	}

	return api.User{
		ID:            userID,
		Name:          "Test user",
		Age:           20,
		Height:        185,
		Sex:           "female",
		ActivityLevel: 5,
		Email:         "test@mail.com",
	}, nil
}

func TestCreateWeightEntry(t *testing.T) {
	mockRepo := mockWeightRepo{}
	mockUserService := api.NewWeightService(&mockRepo)

	tests := []struct {
		name    string
		request api.NewWeightRequest
		want    error
	}{
		{name: "should create a new user successfully", request: api.NewWeightRequest{
			Weight: 70,
			UserID: 1,
		}, want: nil},
		{name: "should return a error because user already exists", request: api.NewWeightRequest{
			Weight: 70,
			UserID: 2,
		}, want: errors.New("storage - user doesn't exists")},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mockUserService.New(test.request)

			if !reflect.DeepEqual(err, test.want) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}
}

func TestCalculateBMR(t *testing.T) {
	mockRepo := mockWeightRepo{}
	mockUserService := api.NewWeightService(&mockRepo)

	tests := []struct {
		name   string
		Height int
		Age    int
		Weight int
		Sex    string
		want   int
		err    error
	}{
		{
			name:   "should calculate BMR for a female",
			Height: 170,
			Age:    22,
			Weight: 65,
			Sex:    "female",
			want:   1441,
			err:    nil,
		},
		{
			name:   "should calculate BMR for a male",
			Height: 170,
			Age:    22,
			Weight: 65,
			Sex:    "male",
			want:   1607,
			err:    nil,
		},
		{
			name:   "should return error because sex wasn't properly specified",
			Height: 170,
			Age:    22,
			Weight: 65,
			Sex:    "",
			want:   0,
			err:    errors.New("invalid variable sex provided to CalculateBMR. needs to be either male or female"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			BMR, err := mockUserService.CalculateBMR(test.Height, test.Age, test.Weight, test.Sex)

			if !reflect.DeepEqual(err, test.err) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.want)
			}

			if !reflect.DeepEqual(BMR, test.want) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, BMR, test.want)
			}
		})
	}
}

func TestDailyIntake(t *testing.T) {
	mockRepo := mockWeightRepo{}
	mockUserService := api.NewWeightService(&mockRepo)

	tests := []struct {
		name          string
		BMR           int
		ActivityLevel int
		want          int
		err           error
	}{
		{
			name:          "should calculate daily intake for activity level 1",
			BMR:           1441,
			ActivityLevel: 1,
			want:          1729,
			err:           nil,
		},
		{
			name:          "should calculate daily intake for activity level 2",
			BMR:           1441,
			ActivityLevel: 2,
			want:          1981,
			err:           nil,
		},
		{
			name:          "should calculate daily intake for activity level 3",
			BMR:           1441,
			ActivityLevel: 3,
			want:          2233,
			err:           nil,
		},
		{
			name:          "should calculate daily intake for activity level 4",
			BMR:           1441,
			ActivityLevel: 4,
			want:          2485,
			err:           nil,
		},
		{
			name:          "should calculate daily intake for activity level 5",
			BMR:           1441,
			ActivityLevel: 5,
			want:          2737,
			err:           nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			BMR, err := mockUserService.DailyIntake(test.BMR, test.ActivityLevel)

			if !reflect.DeepEqual(err, test.err) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, err, test.want)
			}

			if !reflect.DeepEqual(BMR, test.want) {
				t.Errorf("test: %v failed. got: %v, wanted: %v", test.name, BMR, test.want)
			}
		})
	}
}
