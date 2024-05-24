package services

import (
	"errors"
	"sync"

	"Assignment_3_Defense/models"
)

type UserService struct {
	mu     sync.Mutex
	users  map[int]*models.User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]*models.User),
		nextID: 1,
	}
}

func (s *UserService) CreateUser(user *models.User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user.ID = s.nextID
	s.users[s.nextID] = user
	s.nextID++
}

func (s *UserService) GetUser(id int) (*models.User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, found := s.users[id]
	return user, found
}

func (s *UserService) UpdateUser(id int, updatedUser *models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	user, found := s.users[id]
	if !found {
		return errors.New("user not found")
	}
	user.ID = id
	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	return nil
}

func (s *UserService) DeleteUser(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, found := s.users[id]; !found {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}
