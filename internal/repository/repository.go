package repository

import (
	"context"
	"github.com/mereiamangeldin/effective-mobile-test/api"
	"github.com/mereiamangeldin/effective-mobile-test/internal/entity"
)

type IPerson interface {
	AddPerson(ctx context.Context, person *entity.Person) error
	UpdatePerson(ctx context.Context, personId uint, person entity.Person) error
	GetPerson(ctx context.Context, personId uint) (entity.Person, error)
	DeletePerson(ctx context.Context, personId uint) error
	GetPeople(ctx context.Context, peopleFilter api.PeopleFilter, pagination api.Pagination) ([]entity.Person, error)
}

type IRepository interface {
	IPerson
}
