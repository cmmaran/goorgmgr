package memory

import (
	"errors"
	"sync"

	"github.com/cmmaran/goorgmgr/api/model"
)

type Employee struct {
	mutex   sync.Mutex
	records []*model.Employee
}

func NewEmployee() *Employee {
	return new(Employee)
}

func (e *Employee) Create(emp *model.Employee) (*model.Employee, error) {

	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.records = append(e.records, emp)
	return emp, nil
}

func (e *Employee) List() ([]*model.Employee, error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	return e.records, nil
}

func (e *Employee) GetByID(id int64) (*model.Employee, error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for _, rec := range e.records {
		if rec.Id == id {
			return rec, nil
		}
	}

	return nil, errors.New("record not found")
}

func (e *Employee) Delete(id int64) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for i, rec := range e.records {
		if rec.Id == id {
			e.records = append(e.records[:i], e.records[i+1:]...)
			return nil
		}
	}

	return errors.New("record not found")
}

func (e *Employee) Update(id int64, emp *model.Employee) (*model.Employee, error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	for i, rec := range e.records {
		if rec.Id == id {
			e.records[i] = emp
			return e.records[i], nil
		}
	}

	return nil, errors.New("record not found")
}
