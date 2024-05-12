package employee

import (
	"sync"

	"github.com/cmmaran/goorgmgr/api/model"
)

type EmployeeInterfacer interface {
	CreateEmployee(*model.Employee) *model.Employee
	GetEmployee() []*model.Employee
	GetEmployeeByID(id int64) *model.Employee
	DeleteEmployee(id int64) []*model.Employee
	UpdateEmployee(id int64, emp *model.Employee) *model.Employee
}

type employee struct {
	mutex   sync.Mutex
	records []*model.Employee
}

func NewEmployee() *employee {
	return &employee{}
}

func (e *employee) CreateEmployee(emp *model.Employee) *model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.records = append(e.records, emp)
	return emp
}

func (e *employee) GetEmployee() []*model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.records
}

func (e *employee) GetEmployeeByID(id int64) *model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for _, rec := range e.records {
		if rec.Id == id {
			return rec
		}
	}

	return nil
}

func (e *employee) DeleteEmployee(id int64) []*model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for i, rec := range e.records {
		if rec.Id == id {
			e.records = append(e.records[:i], e.records[i+1:]...)
		}
	}

	return e.records
}

func (e *employee) UpdateEmployee(id int64, emp *model.Employee) *model.Employee {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for i, rec := range e.records {
		if rec.Id == id {
			e.records[i] = emp
			return e.records[i]
		}
	}

	return nil
}
