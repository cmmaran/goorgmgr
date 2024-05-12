package business

import (
	"github.com/cmmaran/goorgmgr/internal/layer"
	"github.com/cmmaran/goorgmgr/internal/layer/business/employee"
)

type Business struct {
	E employee.EmployeeInterfacer
}

func (b *Business) Employee() employee.EmployeeInterfacer {
	return b.E
}

func NewBusiness() layer.Layer {
	return &Business{
		E: employee.NewEmployee(),
	}
}
