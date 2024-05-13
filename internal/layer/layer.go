package layer

import (
	"github.com/cmmaran/goorgmgr/internal/layer/business/employee"
)

type Layer interface {
	Employee() employee.EmployeeInterface
}
