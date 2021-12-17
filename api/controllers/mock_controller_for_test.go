package controllers

import (
	"PROJECT1/dataservices"
)

type mockControllerDescriber struct {
	Case int
}

func newMockControllerDescriber(c int) mockControllerDescriber {
	return mockControllerDescriber{Case: c}
}

func (m mockControllerDescriber) Getall() (user []dataservices.User, err *dataservices.Error) {
	user = []dataservices.User{

		{1, "steve", "LA", 1234567890, 5.8, true},
		{2, "Sachin", "Uttarakhand", 2345678902, 5.9, false},
		{3, "Virat", "Mumbai", 7778889901, 8.4, true},
	}
	if m.Case == 2 {
		return nil, &dataservices.Error{m.Case, "Error"}
	}
	return user, nil
}
func (m mockControllerDescriber) GetByID(i int) (*dataservices.User, *dataservices.Error) {
	if m.Case == 5 {
		return nil, &dataservices.Error{i, "User does not exist"}
	}
	return &dataservices.User{2, "Sachin", "Uttarakhand", 2345678902, 5.9, false}, nil
}
