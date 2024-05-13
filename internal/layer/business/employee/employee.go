package employee

import (
	"sync"

	"github.com/cmmaran/goorgmgr/api/model"
	"github.com/cmmaran/goorgmgr/internal/layer/db"
)

type EmployeeInterface interface {
	CreateEmployee(*model.Employee) *model.Employee
	GetEmployee() []*model.Employee
	GetEmployeeByID(id int64) *model.Employee
	DeleteEmployee(id int64) []*model.Employee
	UpdateEmployee(id int64, emp *model.Employee) *model.Employee
}

type employee struct {
	mutex sync.Mutex
	db    *db.DB
}

func NewEmployee() *employee {
	return &employee{
		db: db.NewDB(),
	}
}

func (e *employee) CreateEmployee(emp *model.Employee) *model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	rec, _ := e.db.Employee.Create(emp)
	return rec
}

func (e *employee) GetEmployee() []*model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	recs, _ := e.db.Employee.List()
	return recs
}

func (e *employee) GetEmployeeByID(id int64) *model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	rec, _ := e.db.Employee.GetByID(id)
	return rec
}

func (e *employee) DeleteEmployee(id int64) []*model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.db.Employee.Delete(id)
	return nil
}

func (e *employee) UpdateEmployee(id int64, emp *model.Employee) *model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	rec, _ := e.db.Employee.Update(id, emp)
	return rec
}
