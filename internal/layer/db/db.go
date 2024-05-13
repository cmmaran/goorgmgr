package db

import (
	"github.com/cmmaran/goorgmgr/internal/layer/db/memory"
)

type DB struct {
	Employee *memory.Employee
}

func NewDB() *DB {
	return &DB{
		Employee: memory.NewEmployee(),
	}
}
