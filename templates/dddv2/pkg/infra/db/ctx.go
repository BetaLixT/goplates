package db

import (
	"github.com/BetaLixT/tsqlx"
	"github.com/jmoiron/sqlx"
)

func NewDatabaseContext(
  tracer tsqlx.ITracer,
  optn *Options,
) (*tsqlx.TracedDB, error) {
  db, err := sqlx.Open("postgres", optn.ConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

  return tsqlx.NewTracedDB(
    db,
    tracer,
    optn.DatabaseServiceName,
  ), nil
}
