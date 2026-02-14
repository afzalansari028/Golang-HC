package main

import (
	"errors"
	"testing"
)

// implement mock success repo
type MockSuccessRepo struct{}

func (mks *MockSuccessRepo) GetById(id int) (User, error) {

	if id == 10 {
		return User{
			Id:   10,
			Name: "Virat",
		}, nil
	}
	return User{}, errors.New("User not found")

}

// implement mock fail repo
type MockErrorRepo struct{}

func (mkerr *MockErrorRepo) GetById(id int) (User, error) {
	return User{}, errors.New("Database failure")
}

func TestGetUser_Success(t *testing.T) {

	mockRepo := &MockSuccessRepo{}
	service := NewUserService(mockRepo)

	user, err := service.GetUser(10)
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}

	if user.Id != 10 {
		t.Errorf("Expected 10 but got %v", user.Id)
	}

}
func TestGetUser_InvalidID(t *testing.T) {

	mockRepo := &MockSuccessRepo{}
	service := NewUserService(mockRepo)

	_, err := service.GetUser(20)

	if err == nil {
		t.Errorf("expected error for invalid ID:%v", err)
	}
}

func TestGetUser_DBError(t *testing.T) {

	mockRepo := &MockErrorRepo{}
	service := NewUserService(mockRepo)

	_, err := service.GetUser(5)

	if err == nil {
		t.Errorf("expected database error")
	}
}
