package pg

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mereiamangeldin/effective-mobile-test/pkg/logging"
	"net/url"
)

const peopleTable = "people"

type Postgres struct {
	host     string
	username string
	password string
	port     string
	dbName   string
	Pool     *pgxpool.Pool
	SqlDB    *sql.DB
	logger   logging.Logger
}

func New(logger logging.Logger, opts ...Option) (*Postgres, error) {
	p := new(Postgres)
	p.logger = logger

	for _, opt := range opts {
		opt(p)
	}

	q := url.Values{}
	q.Add("sslmode", "disable")

	u := url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword(p.username, p.password),
		Host:     fmt.Sprintf("%s:%s", p.host, p.port),
		Path:     p.dbName,
		RawQuery: q.Encode(),
	}

	poolConfig, err := pgxpool.ParseConfig(u.String())
	if err != nil {
		return nil, fmt.Errorf("pgxpool parse config err: %w", err)
	}

	p.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect err: %w", err)
	}

	return p, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
