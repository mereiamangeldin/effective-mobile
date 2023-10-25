package pg

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/mereiamangeldin/effective-mobile-test/api"
	"github.com/mereiamangeldin/effective-mobile-test/internal/entity"
)

func (p *Postgres) AddPerson(ctx context.Context, person *entity.Person) error {
	p.logger.Info("adding a new person to database")
	query := fmt.Sprintf(`
	INSERT INTO %s (
	                name,
	                surname,
	                patronymic,
	                age,
	                gender,
	                nationality
					)
		VALUES($1, $2, $3, $4, $5, $6)
`, peopleTable)

	_, err := p.Pool.Exec(ctx, query, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality)
	if err != nil {
		return err
	}
	p.logger.Info("person is successfully added to database")
	return nil
}

func (p *Postgres) UpdatePerson(ctx context.Context, personId uint, person entity.Person) error {
	p.logger.Info("updating a person in database")
	query := fmt.Sprintf("UPDATE %s SET  name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, nationality = $6 where id = $7", peopleTable)
	_, err := p.Pool.Exec(ctx, query, person.Name, person.Surname, person.Patronymic, person.Age, person.Gender, person.Nationality, personId)
	if err != nil {
		return err
	}
	p.logger.Info("person is successfully updated in database")
	return nil
}

func (p *Postgres) GetPerson(ctx context.Context, personId uint) (entity.Person, error) {
	p.logger.Info("getting a person from database")
	var person entity.Person
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", peopleTable)
	err := pgxscan.Get(ctx, p.Pool, &person, query, personId)
	if err != nil {
		return entity.Person{}, err
	}
	p.logger.Info("person is successfully returned from database")
	return person, nil
}

func (p *Postgres) DeletePerson(ctx context.Context, personId uint) error {
	p.logger.Info("deleting a person from database")
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", peopleTable)

	_, err := p.Pool.Exec(ctx, query, personId)

	if err != nil {
		return err
	}
	p.logger.Info("person is successfully deleted from database")
	return nil
}

func (p *Postgres) GetPeople(ctx context.Context, peopleFilter api.PeopleFilter, pagination api.Pagination) ([]entity.Person, error) {
	p.logger.Info("getting a people from database")
	var people []entity.Person
	query := ""
	var args []interface{}
	if peopleFilter.Gender != "" {
		query = fmt.Sprintf("SELECT * FROM %s WHERE gender = $1 LIMIT $2 OFFSET $3", peopleTable)
		args = append(args, peopleFilter.Gender)
	} else {
		query = fmt.Sprintf("SELECT * FROM %s LIMIT $1 OFFSET $2", peopleTable)
	}

	args = append(args, pagination.Limit, pagination.Offset)
	err := pgxscan.Select(ctx, p.Pool, &people, query, args...)
	if err != nil {
		return nil, err
	}
	p.logger.Info("people are successfully returned from database")
	return people, nil
}
