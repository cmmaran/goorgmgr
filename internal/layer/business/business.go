package business

import (
	"github.com/cmmaran/goorgmgr/internal/layer"
	"github.com/cmmaran/goorgmgr/internal/layer/business/employee"
)

type Business struct {
	E employee.EmployeeInterface
}

func (b *Business) Employee() employee.EmployeeInterface {
	return b.E
}

func NewBusiness() layer.Layer {
	return &Business{
		E: employee.NewEmployee(),
	}
}
