package service

import (
	"context"
	"github.com/mereiamangeldin/effective-mobile-test/api"
	"github.com/mereiamangeldin/effective-mobile-test/internal/entity"
)

func (m *Manager) AddPerson(ctx context.Context, person *entity.Person) error {
	return m.Repository.AddPerson(ctx, person)
}

func (m *Manager) UpdatePerson(ctx context.Context, personId uint, person entity.Person) error {
	return m.Repository.UpdatePerson(ctx, personId, person)
}

func (m *Manager) GetPerson(ctx context.Context, personId uint) (entity.Person, error) {
	return m.Repository.GetPerson(ctx, personId)
}

func (m *Manager) DeletePerson(ctx context.Context, personId uint) error {
	return m.Repository.DeletePerson(ctx, personId)
}

func (m *Manager) GetPeople(ctx context.Context, peopleFilter api.PeopleFilter, pagination api.Pagination) ([]entity.Person, error) {
	return m.Repository.GetPeople(ctx, peopleFilter, pagination)
}
