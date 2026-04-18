package service

import (
	"fmt"
	"practice-8/repository"
	"testing"
)

func TestRegisterUser_AllCases(t *testing.T) {

	t.Run("user exists", func(t *testing.T) {
		mock := &repository.MockRepo{User: &repository.User{}}
		s := NewUserService(mock)

		err := s.RegisterUser(&repository.User{}, "mail")
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("repo error", func(t *testing.T) {
		mock := &repository.MockRepo{Err: fmt.Errorf("db error")}
		s := NewUserService(mock)

		err := s.RegisterUser(&repository.User{}, "mail")
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("success", func(t *testing.T) {
		mock := &repository.MockRepo{}
		s := NewUserService(mock)

		err := s.RegisterUser(&repository.User{}, "mail")
		if err != nil {
			t.Errorf("unexpected error")
		}
	})
}

func TestUpdateUserName_AllCases(t *testing.T) {

	t.Run("empty name", func(t *testing.T) {
		mock := &repository.MockRepo{}
		s := NewUserService(mock)

		err := s.UpdateUserName(1, "")
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("user not found", func(t *testing.T) {
		mock := &repository.MockRepo{Err: fmt.Errorf("not found")}
		s := NewUserService(mock)

		err := s.UpdateUserName(1, "New")
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("update fails", func(t *testing.T) {
		mock := &repository.MockRepo{
			User: &repository.User{},
			Err:  fmt.Errorf("update error"),
		}
		s := NewUserService(mock)

		err := s.UpdateUserName(1, "New")
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("success and verify name changed", func(t *testing.T) {
		user := &repository.User{Name: "Old"}
		mock := &repository.MockRepo{User: user}
		s := NewUserService(mock)

		_ = s.UpdateUserName(1, "New")

		if user.Name != "New" {
			t.Errorf("name was not updated")
		}
	})
}

func TestDeleteUser_AllCases(t *testing.T) {

	t.Run("admin delete", func(t *testing.T) {
		mock := &repository.MockRepo{}
		s := NewUserService(mock)

		err := s.DeleteUser(1)
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("repo error", func(t *testing.T) {
		mock := &repository.MockRepo{Err: fmt.Errorf("fail")}
		s := NewUserService(mock)

		err := s.DeleteUser(2)
		if err == nil {
			t.Errorf("expected error")
		}
	})

	t.Run("success", func(t *testing.T) {
		mock := &repository.MockRepo{}
		s := NewUserService(mock)

		err := s.DeleteUser(2)
		if err != nil {
			t.Errorf("unexpected error")
		}
	})
}
