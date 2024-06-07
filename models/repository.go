package models

import "database/sql"

// Repository is the db repo.  Anything that implements
// this interface must implement all the methods here.
type Repository interface {
    AllDogBreeds() ([]*DogBreed, error)
}

// mysqlRepository is a simple wrapper for the *sql.DB type.
// This is used to return a MySQL/MariaDB repo.
type mysqlRepository struct {
    DB *sql.DB
}

// newMysqlRepository is convenience factory to return
// a new newMysqlRepository.
func newMysqlRepository(conn *sql.DB) Repository {
    return &mysqlRepository {
        DB: conn,
    }
}

type testRepository struct {
    DB *sql.DB
}

func newTestRepository(conn *sql.DB) Repository {
    return &testRepository {
        DB: nil,
    }
}

