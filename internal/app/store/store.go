package store

import (
	"database/sql"
	_ "github.com/lib/pq" // ...
)

type Store struct {
	config             *Config
	db                 *sql.DB
	recorderRepository *RecorderRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Record() *RecorderRepository {
	if s.recorderRepository != nil {
		return s.recorderRepository
	}

	s.recorderRepository = &RecorderRepository{
		store: s,
	}
	return s.recorderRepository
}
