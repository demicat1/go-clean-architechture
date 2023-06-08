package usecase

import (
	"adilkhan.go-clean-architechture/services/contact/internal/domain"
	"adilkhan.go-clean-architechture/services/contact/internal/repository"
)

type UseCase interface {
	CreateContact(user *domain.Contact) (int, error)
	ReadContact(id int) (*domain.Contact, error)
	UpdateContact(id int) error
	DeleteContact(id int) error
	CreateGroup(group *domain.Group) (int, error)
	ReadGroup(id int) (*domain.Group, error)
	AddContact(id int) error
}

type useCase struct {
	repo repository.Repo
}

func New(repo repository.Repo) *useCase {
	return &useCase{repo: repo}
}
